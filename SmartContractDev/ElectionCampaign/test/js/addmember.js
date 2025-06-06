// 确保路径正确
import { CONTRACT_ABI, CONTRACT_ADDRESS } from './contractConfig.js';

let provider;
let signer;
let contract;

document.addEventListener('DOMContentLoaded', async () => {
    if (typeof window.ethereum !== 'undefined') {
        try {
            await window.ethereum.request({ method: 'eth_requestAccounts' });
            provider = new ethers.BrowserProvider(window.ethereum);
            signer = await provider.getSigner();
            contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer);
        } catch (error) {
            console.error('连接 Metamask 失败:', error);
            alert('请先连接 Metamask');
        }
    } else {
        alert('请安装 Metamask 插件');
    }

    const goToVoteBtn = document.getElementById('goToVoteBtn');
    const submitBtn = document.getElementById('submitBtn');
    // 直接使用全局的 bootstrap 对象
    const submitResultModal = new bootstrap.Modal(document.getElementById('submitResultModal'));

    goToVoteBtn.addEventListener('click', () => {
        window.location.href = '../pages/voting.html';
    });
    submitBtn.addEventListener('click', async () => {
        await registerCandidate(submitResultModal);
    });
});

async function registerCandidate(submitResultModal) {
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
        document.getElementById('candidateName').value = '';
        document.getElementById('candidateDescription').value = '';
        submitResultModal.show();
    } catch (error) {
        console.error('参赛提交失败:', error);
        alert('参赛提交失败');
    }
}
