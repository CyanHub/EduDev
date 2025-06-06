import { CONTRACT_ABI, CONTRACT_ADDRESS } from './contractConfig.js';

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

    setupEventListeners();
    await loadChart('bar');
});

function setupEventListeners() {
    setupChartTypeButtons();
    setupNavigationButtons();
}

function setupChartTypeButtons() {
    const barChartBtn = document.getElementById('barChartBtn');
    const pieChartBtn = document.getElementById('pieChartBtn');

    // 输出获取到的按钮元素，方便确认是否成功获取
    console.log('barChartBtn:', barChartBtn);
    console.log('pieChartBtn:', pieChartBtn);

    if (!barChartBtn || !pieChartBtn) {
        console.error('图表类型按钮获取失败，请检查按钮 ID 和 HTML 结构');
        return;
    }

    barChartBtn.addEventListener('click', () => loadChart('bar'));
    pieChartBtn.addEventListener('click', () => loadChart('pie'));
}

function setupNavigationButtons() {
    const goToVoteBtn = document.getElementById('goToVoteBtn');
    const joinCampaignBtn = document.getElementById('joinCampaignBtn');

    // 输出获取到的按钮元素，方便确认是否成功获取
    console.log('goToVoteBtn:', goToVoteBtn);
    console.log('joinCampaignBtn:', joinCampaignBtn);

    if (!goToVoteBtn || !joinCampaignBtn) {
        console.error('导航按钮获取失败，请检查按钮 ID 和 HTML 结构');
        return;
    }

    goToVoteBtn.addEventListener('click', () => {
            window.location.href = '../pages/voting.html';
    });

    joinCampaignBtn.addEventListener('click', () => {
            window.location.href = '../pages/addmember.html';
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

async function loadChart(type) {
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

        // 按票数降序排序
        candidates.sort((a, b) => b.voteCount - a.voteCount);

        const names = candidates.map(c => c.name);
        const voteCounts = candidates.map(c => c.voteCount);

        const chartDom = document.getElementById('chartContainer');
        if (!myChart) {
            myChart = echarts.init(chartDom);
        } else {
            // 清除当前图表
            myChart.clear();
        }

        let option;
        if (type === 'bar') {
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
            const totalVotes = voteCounts.reduce((sum, count) => sum + count, 0);
            option = {
                // 饼图不需要 xAxis 和 yAxis，移除这两个配置
                series: [
                    {
                        type: 'pie',
                        data: candidates.map(c => ({
                            value: c.voteCount,
                            name: `${c.name} (${((c.voteCount / totalVotes) * 100).toFixed(2)}%)`
                        })),
                        label: {
                            formatter: '{b}: {d}%'
                        }
                    }
                ]
            };
        }

        myChart.setOption(option);
    } catch (error) {
        console.error('加载图表失败:', error);
        alert('加载图表失败');
    }
}
