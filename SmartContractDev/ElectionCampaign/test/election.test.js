const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Election Contract", function () {
  let Election;
  let election;
  let owner;
  let addr1;

  beforeEach(async function () {
    [owner, addr1] = await ethers.getSigners();
    Election = await ethers.getContractFactory("Election");
    election = await Election.deploy();
    await election.waitForDeployment();
  });

  it("Should register a candidate", async function () {
    await election.registerCandidate(addr1.address, "Candidate 1");
    const candidate = await election.candidates(0);
    expect(candidate.name).to.equal("Candidate 1");
  });

  it("Should allow voting", async function () {
    await election.registerCandidate(addr1.address, "Candidate 1");
    await election.vote(0);
    const votes = await election.getCandidateVoteCount(0);
    expect(votes).to.equal(1);
  });
});
