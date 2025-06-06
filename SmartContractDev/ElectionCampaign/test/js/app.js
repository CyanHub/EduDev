import { CONTRACT_ABI, CONTRACT_ADDRESS } from './contractConfig.js';  // 若路径有变化需对应修改

let provider;
let signer;
let contract;
let currentPage = 1;

document.addEventListener('DOMContentLoaded', () => {
    const connectBtn = document.getElementById('connectBtn');
    const confirmBtn = document.getElementById('confirmBtn');
    const successModal = document.getElementById('successModal');

    if (!connectBtn || !confirmBtn || !successModal) {
        console.error('某些元素未找到，请检查 HTML 中的 ID 是否正确。');
        return;
    }

    connectBtn.addEventListener('click', async () => {
        await connectMetamask();
        if (contract) {
            successModal.style.display = 'flex';
        }
    });

    confirmBtn.addEventListener('click', () => {
        successModal.style.display = 'none';
        // 跳转到投票页面
        window.location.href = 'voting.html';
    });
});

async function connectMetamask() {
    if (typeof window.ethereum !== 'undefined') {
        try {
            // 请求用户授权账户
            await window.ethereum.request({ method: 'eth_requestAccounts' });
            provider = new ethers.BrowserProvider(window.ethereum);
            signer = await provider.getSigner();
            contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer);
        } catch (error) {
            console.error('连接 Metamask 失败:', error);
            alert('连接 Metamask 失败');
        }
    } else {
        alert('请安装 Metamask 插件');
    }
}

async function loadCandidates() {
    if (!contract) return;

    try {
        const candidatesCount = await contract.getCandidatesCount();
        const itemsPerPage = parseInt(document.getElementById('itemsPerPage').value);
        const startIndex = (currentPage - 1) * itemsPerPage;
        const endIndex = Math.min(startIndex + itemsPerPage, candidatesCount);

        const tableBody = document.querySelector('#candidateTable tbody');
        tableBody.innerHTML = '';

        for (let i = startIndex; i < endIndex; i++) {
            const candidate = await contract.getCandidate(i);
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${candidate[1]}</td>
                <td>${candidate[3]}</td>
                <td>${candidate[2]}</td>
                <td><button onclick="voteForCandidate(${i})">投票</button></td>
            `;
            tableBody.appendChild(row);
        }
    } catch (error) {
        console.error('加载候选人信息失败:', error);
    }
}

async function voteForCandidate(index) {
    if (!contract) return;

    try {
        const tx = await contract.vote(index);
        await tx.wait();
        alert('投票成功');
        loadCandidates();
    } catch (error) {
        console.error('投票失败:', error);
        alert('投票失败');
    }
}

async function viewRank() {
    if (!contract) return;

    try {
        const candidatesCount = await contract.getCandidatesCount();
        const candidates = [];

        for (let i = 0; i < candidatesCount; i++) {
            const candidate = await contract.getCandidate(i);
            candidates.push({
                index: i,
                name: candidate[1],
                voteCount: candidate[2]
            });
        }

        // 按票数降序排序
        candidates.sort((a, b) => b.voteCount - a.voteCount);

        const rankInfo = candidates.map((candidate, index) => {
            return `${index + 1}. ${candidate.name}: ${candidate.voteCount} 票`;
        }).join('\n');

        alert(`候选人排行：\n${rankInfo}`);
    } catch (error) {
        console.error('查看排行失败:', error);
        alert('查看排行失败');
    }
}

function prevPage() {
    if (currentPage > 1) {
        currentPage--;
        document.getElementById('pageNumber').value = currentPage;
        loadCandidates();
    }
}

function goToPage() {
    const pageNumber = parseInt(document.getElementById('pageNumber').value);
    if (pageNumber > 0) {
        currentPage = pageNumber;
        loadCandidates();
    }
}

function nextPage() {
    // 这里简化处理，实际应根据候选人总数和每页显示数量计算最大页数
    currentPage++;
    document.getElementById('pageNumber').value = currentPage;
    loadCandidates();
}

async function registerCandidate() {
    if (!contract) return;

    const name = document.getElementById('candidateName').value;
    const description = document.getElementById('candidateDescription').value;

    if (!name || !description) {
        alert('请输入候选人姓名和简介');
        return;
    }

    try {
        const signerAddress = await signer.getAddress();
        const tx = await contract.registerCandidate(signerAddress, name, description);
        await tx.wait();
        alert('参选提交成功');
        document.getElementById('candidateName').value = '';
        document.getElementById('candidateDescription').value = '';
        loadCandidates();
    } catch (error) {
        console.error('参选提交失败:', error);
        alert('参选提交失败');
    }
}

// 暴露函数到全局作用域，以便 HTML 中的 onclick 调用
window.voteForCandidate = voteForCandidate;