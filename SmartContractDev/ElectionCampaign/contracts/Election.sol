// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Election {
    // 候选人结构体
    struct Candidate {
        address candidateAddress;
        string name;
        uint256 voteCount;
        string description; // 新增候选人简介
    }

    // 候选人数组
    Candidate[] public candidates;

    // 新增获取候选人数量的方法
    function getCandidatesCount() external view returns (uint256) {
        return candidates.length;
    }

    // 获取候选人信息的方法
    function getCandidate(
        uint256 _candidateIndex
    ) external view returns (address, string memory, uint256, string memory) {
        require(_candidateIndex < candidates.length, "Invalid candidate index");
        Candidate memory candidate = candidates[_candidateIndex];
        return (
            candidate.candidateAddress,
            candidate.name,
            candidate.voteCount,
            candidate.description
        );
    }

    // 选民地址到是否已投票的映射
    mapping(address => bool) public voters;

    // 构造函数，添加初始候选人
    constructor() {
        // 添加初始候选人，地址使用部署者地址，实际可按需修改
        // registerCandidate(msg.sender, "林晚荣", "天下第一丁");  // 这里在终端提示不符合Unicode编码，暂时注释掉
    }

    // 注册竞选人，新增简介参数
    function registerCandidate(
        address _candidateAddress,
        string memory _name,
        string memory _description
    ) external {
        candidates.push(Candidate(_candidateAddress, _name, 0, _description));
    }

    // 投票
    function vote(uint256 _candidateIndex) external {
        require(!voters[msg.sender], "You have already voted.");
        require(
            _candidateIndex < candidates.length,
            "Invalid candidate index."
        );

        candidates[_candidateIndex].voteCount++;
        voters[msg.sender] = true;
    }

    // 查询候选人得票数
    function getCandidateVoteCount(
        uint256 _candidateIndex
    ) external view returns (uint256) {
        require(
            _candidateIndex < candidates.length,
            "Invalid candidate index."
        );
        return candidates[_candidateIndex].voteCount;
    }
}
