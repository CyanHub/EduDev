// 用于实现与 Metamask 连接和绘制 Echarts 图表的逻辑。
// 确保在页面加载完成后，DOMContentLoaded 事件被触发，然后执行 connectMetamask 和 loadCampaignChart 函数。
// 确保在 connectMetamask 函数中，使用 ethers.BrowserProvider 来获取用户的账户信息，并使用 ethers.Contract 来实例化智能合约。

// 目前已废弃，后续会进行更新。
import { CONTRACT_ABI, CONTRACT_ADDRESS } from './contractConfig.js';  // 若路径有变化需对应修改
let provider;
let signer;
let contract;
let myChart;

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

    const loadChartBtn = document.getElementById('loadChartBtn');
    const chartTypeSelect = document.getElementById('chartType');

    loadChartBtn.addEventListener('click', loadCampaignChart);
    chartTypeSelect.addEventListener('change', loadCampaignChart);

    await loadCandidates();
});

async function loadCampaignChart() {
    if (!contract) {
        alert('请先连接 Metamask');
        return;
    }

    try {
        const candidatesCount = await contract.getCandidatesCount();
        const candidates = [];

        for (let i = 0; i < candidatesCount; i++) {
            const candidate = await contract.getCandidate(i);
            candidates.push({
                name: candidate[1],
                voteCount: parseInt(candidate[2])
            });
        }

        const names = candidates.map(c => c.name);
        const voteCounts = candidates.map(c => c.voteCount);

        const chartDom = document.getElementById('chartContainer');
        if (!myChart) {
            myChart = echarts.init(chartDom);
        }

        const chartType = document.getElementById('chartType').value;
        let option;

        if (chartType === 'bar') {
            option = {
                xAxis: {
                    type: 'category',
                    data: names
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: voteCounts,
                        type: 'bar'
                    }
                ]
            };
        } else {
            option = {
                series: [
                    {
                        type: 'pie',
                        data: candidates.map((c, index) => ({
                            value: c.voteCount,
                            name: c.name
                        }))
                    }
                ]
            };
        }

        myChart.setOption(option);
    } catch (error) {
        console.error('加载竞选统计图表失败:', error);
        alert('加载竞选统计图表失败');
    }
}

async function loadCandidates() {
    if (!contract) return;

    try {
        const candidatesCount = await contract.getCandidatesCount();
        const tableBody = document.querySelector('#candidateTable tbody');
        tableBody.innerHTML = '';

        for (let i = 0; i < candidatesCount; i++) {
            const candidate = await contract.getCandidate(i);
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${candidate[1]}</td>
                <td>${candidate[3]}</td>
                <td>${candidate[2]}</td>
                <td><button class="btn btn-primary" onclick="voteForCandidate(${i})">投票</button></td>
            `;
            tableBody.appendChild(row);
        }

        window.voteForCandidate = voteForCandidate;
    } catch (error) {
        console.error('加载候选人信息失败:', error);
    }
}

async function voteForCandidate(index) {
    if (!contract) return;

    try {
        const hasVoted = await contract.voters(await signer.getAddress());
        if (hasVoted) {
            alert('你已经投过票了，每个用户只能投一票。');
            return;
        }

        const tx = await contract.vote(index);
        await tx.wait();
        alert('投票成功');
        await loadCandidates();
        await loadCampaignChart();
    } catch (error) {
        console.error('投票失败:', error);
        alert('投票失败');
    }
}
