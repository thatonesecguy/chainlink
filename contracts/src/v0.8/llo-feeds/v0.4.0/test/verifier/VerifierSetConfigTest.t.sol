// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {BaseTest, BaseTestWithMultipleConfiguredDigests} from "./BaseDestinationVerifierTest.t.sol";
import {DestinationVerifier} from "../../../v0.4.0/DestinationVerifier.sol";
import {Common} from "../../../libraries/Common.sol";

contract VerifierSetConfigTest is BaseTest {
  function setUp() public virtual override {
    BaseTest.setUp();
  }

  function test_revertsIfCalledByNonOwner() public {
    vm.expectRevert("Only callable by owner");
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    changePrank(USER);
    s_verifier.setConfig(
      _getSignerAddresses(signers),
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_revertsIfSetWithTooManySigners() public {
    address[] memory signers = new address[](MAX_ORACLES + 1);
    vm.expectRevert(abi.encodeWithSelector(DestinationVerifier.ExcessSigners.selector, signers.length, MAX_ORACLES));
    s_verifier.setConfig(
      signers,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_revertsIfFaultToleranceIsZero() public {
    vm.expectRevert(abi.encodeWithSelector(DestinationVerifier.FaultToleranceMustBePositive.selector));
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    s_verifier.setConfig(
      _getSignerAddresses(signers),
      0,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_revertsIfNotEnoughSigners() public {
    address[] memory signers = new address[](2);
    signers[0] = address(1000);
    signers[1] = address(1001);

    vm.expectRevert(
      abi.encodeWithSelector(DestinationVerifier.InsufficientSigners.selector, signers.length, FAULT_TOLERANCE * 3 + 1)
    );
    s_verifier.setConfig(
      signers,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_revertsIfDuplicateSigners() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    address[] memory signerAddrs = _getSignerAddresses(signers);
    signerAddrs[0] = signerAddrs[1];
    vm.expectRevert(abi.encodeWithSelector(DestinationVerifier.NonUniqueSignatures.selector));
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_revertsIfSignerContainsZeroAddress() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    address[] memory signerAddrs = _getSignerAddresses(signers);
    signerAddrs[0] = address(0);
    vm.expectRevert(abi.encodeWithSelector(DestinationVerifier.ZeroAddress.selector));
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_DONConfigIDIsSameForSignersInDifferentOrder() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    address[] memory signerAddrs = _getSignerAddresses(signers);

//    bytes24 expectedDonConfig = 0x63eab508c9125e9cf2b0937afa833ae0c6f371729aa671bd;

    bytes24 expectedDonConfigID = _DONConfigIdFromConfigData(signerAddrs ,FAULT_TOLERANCE );

    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );


   address temp = signerAddrs[0];
   signerAddrs[0] = signerAddrs[1];
   signerAddrs[1] = temp;

 vm.expectRevert(abi.encodeWithSelector(DestinationVerifier.DONConfigAlreadyExists.selector, expectedDonConfigID));
 
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );
  }

  function test_NoDonConfigAlreadyExists() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    address[] memory signerAddrs = _getSignerAddresses(signers);

    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );

   // Testing adding same set of Signers but different FAULT_TOLERENCE does not result in DONConfigAlreadyExists revert
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE - 1,
      new Common.AddressAndWeight[](0)
    );

    // Testing adding a different set of Signers with same FAULT_TOLERENCE does not result in DONConfigAlreadyExists revert
    address[] memory signerAddrsMinusOne = new address[](signerAddrs.length-1);
    for (uint i = 0; i < signerAddrs.length - 1; i++) {
            signerAddrsMinusOne[i] = signerAddrs[i];
        }
    s_verifier.setConfig(
      signerAddrsMinusOne,
      FAULT_TOLERANCE - 1,
      new Common.AddressAndWeight[](0)
    );
  }

function test_addressesAndWeightsDoNotProduceSideEffectsInDonConfigIds() public {

   Signer[] memory signers = _getSigners(MAX_ORACLES);
    address[] memory signerAddrs = _getSignerAddresses(signers);

    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );


  bytes24 expectedDonConfigID = _DONConfigIdFromConfigData(signerAddrs ,FAULT_TOLERANCE );

  //bytes24 expectedDonConfig = 0x63eab508c9125e9cf2b0937afa833ae0c6f371729aa671bd;

 vm.expectRevert(abi.encodeWithSelector(DestinationVerifier.DONConfigAlreadyExists.selector, expectedDonConfigID));

    // Same call to setConfig with different addressAndWeights do not entail a new DONConfigId 
    // Resulting in a DONConfigAlreadyExists error
    Common.AddressAndWeight[] memory weights = new Common.AddressAndWeight[](1);
    weights[0] = Common.AddressAndWeight(signers[0].signerAddress, 1);
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      weights
    );
}

/*
function test_setConfigWithAddressesAndWeightsAreSetCorrectly() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);
    address[] memory signerAddrs = _getSignerAddresses(signers);
    Common.AddressAndWeight[] memory weights = new Common.AddressAndWeight[](1);
    weights[0] = Common.AddressAndWeight(signers[0].signerAddress, 1);
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      weights
    );

   // check internal state of feeManager
   // we nest a BaseTestWithConfiguredVerifierAndFeeManager for this to work
}
*/

// mine
 function test_correctlyUpdatesTheConfig() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);
     address[] memory signerAddrs = _getSignerAddresses(signers);
    
    s_verifier.setConfig(
      signerAddrs,
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );

bytes24 expectedDonConfigID = _DONConfigIdFromConfigData(signerAddrs,FAULT_TOLERANCE );

    //bytes24 expectedDonConfigID = 0x63eab508c9125e9cf2b0937afa833ae0c6f371729aa671bd;

    DestinationVerifier.DONConfig memory generatedDonConfig = s_verifier.getDONConfig(expectedDonConfigID);
    assertEq(generatedDonConfig.f, FAULT_TOLERANCE);
    assertEq(generatedDonConfig.isActive, true);
    assertEq(generatedDonConfig.DONConfigID, expectedDonConfigID);
}

/*

  function test_correctlyUpdatesTheConfig() public {
    Signer[] memory signers = _getSigners(MAX_ORACLES);

    s_verifierProxy.initializeVerifier(address(s_verifier));
    s_verifier.setConfig(
      _getSignerAddresses(signers),
      FAULT_TOLERANCE,
      new Common.AddressAndWeight[](0)
    );



    bytes32 expectedConfigDigest = _configDigestFromConfigData(
      FEED_ID,
      block.chainid,
      address(s_verifier),
      1,
      _getSignerAddresses(signers),
      s_offchaintransmitters,
      FAULT_TOLERANCE,
      bytes(""),
      VERIFIER_VERSION,
      bytes("")
    );

    (uint32 configCount, uint32 blockNumber, bytes32 configDigest) = s_verifier.latestConfigDetails(FEED_ID);
    assertEq(configCount, 1);
    assertEq(blockNumber, block.number);
    assertEq(configDigest, expectedConfigDigest);

    (bool scanLogs, bytes32 configDigestTwo, uint32 epoch) = s_verifier.latestConfigDigestAndEpoch(FEED_ID);
    assertEq(scanLogs, false);
    assertEq(configDigestTwo, expectedConfigDigest);
    assertEq(epoch, 0);
  }
}



contract VerifierSetConfigWhenThereAreMultipleDigestsTest is BaseTestWithMultipleConfiguredDigests {
  function test_correctlyUpdatesTheDigestInTheProxy() public {
    Signer[] memory newSigners = _getSigners(15);

    s_verifier.setConfig(
      FEED_ID,
      _getSignerAddresses(newSigners),
      s_offchaintransmitters,
      4,
      bytes(""),
      VERIFIER_VERSION,
      bytes(""),
      new Common.AddressAndWeight[](0)
    );

    (, , bytes32 configDigest) = s_verifier.latestConfigDetails(FEED_ID);
    address verifierAddr = s_verifierProxy.getVerifier(configDigest);
    assertEq(verifierAddr, address(s_verifier));
  }

  function test_correctlyUpdatesDigestsOnMultipleVerifiersInTheProxy() public {
    Signer[] memory newSigners = _getSigners(15);

    s_verifier.setConfig(
      FEED_ID_2,
      _getSignerAddresses(newSigners),
      s_offchaintransmitters,
      4,
      bytes(""),
      VERIFIER_VERSION,
      bytes(""),
      new Common.AddressAndWeight[](0)
    );

    (, , bytes32 configDigest) = s_verifier.latestConfigDetails(FEED_ID_2);
    address verifierAddr = s_verifierProxy.getVerifier(configDigest);
    assertEq(verifierAddr, address(s_verifier));

    s_verifier_2.setConfig(
      FEED_ID_3,
      _getSignerAddresses(newSigners),
      s_offchaintransmitters,
      4,
      bytes(""),
      VERIFIER_VERSION,
      bytes(""),
      new Common.AddressAndWeight[](0)
    );

    (, , bytes32 configDigest2) = s_verifier_2.latestConfigDetails(FEED_ID_3);
    address verifierAddr2 = s_verifierProxy.getVerifier(configDigest2);
    assertEq(verifierAddr2, address(s_verifier_2));
  }

  function test_correctlySetsConfigWhenDigestsAreRemoved() public {
    s_verifier.deactivateConfig(FEED_ID, s_configDigestTwo);

    Signer[] memory newSigners = _getSigners(15);

    s_verifier.setConfig(
      FEED_ID,
      _getSignerAddresses(newSigners),
      s_offchaintransmitters,
      4,
      bytes(""),
      VERIFIER_VERSION,
      bytes(""),
      new Common.AddressAndWeight[](0)
    );

    bytes32 expectedConfigDigest = _configDigestFromConfigData(
      FEED_ID,
      block.chainid,
      address(s_verifier),
      s_numConfigsSet + 1,
      _getSignerAddresses(newSigners),
      s_offchaintransmitters,
      4,
      bytes(""),
      VERIFIER_VERSION,
      bytes("")
    );

    (uint32 configCount, uint32 blockNumber, bytes32 configDigest) = s_verifier.latestConfigDetails(FEED_ID);

    assertEq(configCount, s_numConfigsSet + 1);
    assertEq(blockNumber, block.number);
    assertEq(configDigest, expectedConfigDigest);
  }
*/

}
