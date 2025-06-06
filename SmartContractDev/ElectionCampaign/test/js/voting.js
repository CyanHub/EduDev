// 确保路径正确
import { CONTRACT_ABI, CONTRACT_ADDRESS } from './contractConfig.js';

let provider;
let signer;
let contract;
let currentPage = 1;
let itemsPerPage = 1;

document.addEventListener('DOMContentLoaded', async () => {
    try {
        if (typeof window.ethereum === 'undefined') {
            alert('请安装 Metamask 插件');
            return;
        }

        await window.ethereum.request({ method: 'eth_requestAccounts' });
        provider = new ethers.BrowserProvider(window.ethereum);
        signer = await provider.getSigner();
        contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer);

        // 检查合约是否成功创建
        if (!contract) {
            console.error('合约创建失败');
            alert('合约创建失败，请检查配置');
            return;
        }

        // 页面加载时调用加载候选人列表函数
        await loadCandidates();

        setupEventListeners();
    } catch (error) {
        console.error('初始化出错:', error);
        alert('初始化出错，请检查 Metamask 连接和配置');
    }
});

function setupEventListeners() {
    const viewRankBtn = document.getElementById('viewRankBtn');
    const joinCampaignBtn = document.getElementById('joinCampaignBtn');
    const backToHomeBtn = document.getElementById('backToHomeBtn');
    const itemsPerPageSelect = document.getElementById('itemsPerPage');
    const goToPageBtn = document.getElementById('goToPageBtn');

    viewRankBtn.addEventListener('click', () => {
        window.location.href = '../pages/ranklist.html';
    });
    joinCampaignBtn.addEventListener('click', () => {
        window.location.href = '../pages/addmember.html';
    });
    backToHomeBtn.addEventListener('click', () => {
        window.location.href = '../pages/index.html';
    });
    itemsPerPageSelect.addEventListener('change', async () => {
        itemsPerPage = parseInt(itemsPerPageSelect.value);
        await loadCandidates();
    });
    goToPageBtn.addEventListener('click', async () => {
        const pageNumber = parseInt(document.getElementById('pageNumber').value);
        if (pageNumber > 0) {
            currentPage = pageNumber;
            await loadCandidates();
        }
    });
}

async function loadCandidates() {
    if (!contract) {
        console.error('合约未初始化');
        return;
    }

    try {
        // 获取候选人总数
        const candidatesCountBN = await contract.getCandidatesCount();
        console.log('获取到的候选人总数 BigInt:', candidatesCountBN);

        // 使用 Number() 转换 BigInt 为 Number 类型
        const candidatesCount = Number(candidatesCountBN);
        console.log('转换后的候选人总数:', candidatesCount);

        const totalPages = Math.ceil(candidatesCount / itemsPerPage);
        const startIndex = (currentPage - 1) * itemsPerPage;
        const endIndex = Math.min(startIndex + itemsPerPage, candidatesCount);

        const tableBody = document.querySelector('#candidateTable tbody');
        tableBody.innerHTML = '';

        for (let i = startIndex; i < endIndex; i++) {
            const candidate = await contract.getCandidate(i);
            console.log(`获取到候选人 ${i} 的原始信息:`, candidate);

            // 检查 candidate 每个元素的类型
            candidate.forEach((item, index) => {
                console.log(`候选人 ${i} 的第 ${index} 个元素:`, item, typeof item);
            });

            // 将所有可能是 BigInt 的数据都转换为 Number 类型
            const voteCount = Number(candidate[2]);

            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${candidate[1]}</td>
                <td>${candidate[3]}</td>
                <td>${voteCount}</td>
                <td><button class="btn btn-primary" onclick="voteForCandidate(${i})">投票</button></td>
            `;
            tableBody.appendChild(row);
        }

        document.getElementById('pageInfo').textContent = `当前第 ${currentPage} 页，共 ${totalPages} 页`;
        window.voteForCandidate = voteForCandidate;
    } catch (error) {
        console.error('加载候选人信息失败:', error);
        alert('加载候选人信息失败，请检查控制台日志');
    }
}

async function voteForCandidate(index) {
    if (!contract) return;

    try {
        const candidate = await contract.getCandidate(index);
        const signerAddress = await signer.getAddress();
        if (candidate[0].toLowerCase() === signerAddress.toLowerCase()) {
            alert('不能投自己');
            return;
        }

        const hasVoted = await contract.voters(signerAddress);
        if (hasVoted) {
            alert('你已经投过票了，每个用户只能投一票。');
            return;
        }

        const tx = await contract.vote(index);
        await tx.wait();
        alert('投票成功');
        await loadCandidates();
    } catch (error) {
        console.error('投票失败:', error);
        alert('投票失败');
    }
}
