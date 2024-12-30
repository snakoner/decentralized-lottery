import React from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Card } from "@/components/ui/card";
import { Loader2, AlertCircle } from "lucide-react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { useState } from "react";
import { TextField, Box } from '@mui/material';
import { Add, Remove } from '@mui/icons-material';
import { ethers } from "ethers";
import { localStorageWalletConnectHandler, CONTRACT_ADDRESS, CONTRACT_ABI } from "../constants";

interface EntryDialogProps {
  isOpen?: boolean;
  onClose?: () => void;
  onConfirm?: () => void;
  isProcessing?: boolean;
  error?: string;
  ethersToSpend?:string;
}

interface EntrySectionProps {
  poolAmount?: string;
  ticketPrice?: string;
  ticketsPerAccount?: number;
  isProcessing?: boolean;
  error?: string;
  onEntrySubmit?: () => void;
}

const QuantitySelector = ({ quantity, setQuantity, initialQuantity = 1, minQuantity = 1, maxQuantity = 100 }) => {
  const handleDecrement = () => {
    setQuantity((prevQuantity) => Math.max(prevQuantity - 1, minQuantity));
  };

  const handleIncrement = () => {
    setQuantity((prevQuantity) => Math.min(prevQuantity + 1, maxQuantity));
  };

  return (
    <Box display="flex" alignItems="center">
      <Button
        variant="default"
        color="primary"
        onClick={handleDecrement}
        disabled={quantity <= minQuantity}
      >
        <Remove />
      </Button>
      <TextField
        value={quantity}
        inputProps={{ readOnly: true, style: { textAlign: 'center' } }}
        variant="outlined"
        size="small"
        id="tickets-num-field"
        style={{ 
          width: '50px', 
          margin: '0 10px',
          backgroundColor: 'white',
          borderRadius: "10px",
          outline: "none",
        }}
      />
      <Button
        variant="default"
        color="primary"
        onClick={handleIncrement}
        disabled={quantity >= maxQuantity}
      >
        <Add />
      </Button>
    </Box>
  );
};


const EntryDialog = ({
  isOpen = true,
  onClose = () => {},
  onConfirm = () => {},
  isProcessing = false,
  error = "",
  ethersToSpend
}: EntryDialogProps) => {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>
			{localStorageWalletConnectHandler() ? 'Confirm Lottery Entry' : 'Connect your wallet to buy tickets'}
			</DialogTitle>
          <DialogDescription>
			{localStorageWalletConnectHandler() ? `You are about to enter the lottery and gonna to spend about ${ethersToSpend} ETH. This action cannot be undone.` : ''}
          </DialogDescription>
        </DialogHeader>

        {error && (
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertDescription>{error}</AlertDescription>
          </Alert>
        )}

        <DialogFooter>
          <Button variant="outline" onClick={onClose} disabled={isProcessing}>
            Cancel
          </Button>
          <Button onClick={onConfirm} disabled={isProcessing}>
            {isProcessing ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Processing
              </>
            ) : (
              "Confirm Entry"
            )}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

const EntrySection = ({
	poolAmount = "100 ETH",
	ticketPrice = "0.1 ETH",
	ticketsPerAccount = 0,
	isProcessing = false,
	error = "",
	onEntrySubmit = () => {},
	}: EntrySectionProps) => {
	const [showDialog, setShowDialog] = React.useState(false);
	const [quantity, setQuantity] = React.useState(1);
	const [ethersToSpend, setEthersToSpend] = React.useState("");

	const buyTickets = async() => {
		if (!localStorageWalletConnectHandler()) {
			alert("Connect your wallet");
			return;
		}

		const txParams = {
			value: ethers.parseEther(ethersToSpend),
		};


		const browserProvider = new ethers.BrowserProvider(window.ethereum);
		const signer = await browserProvider.getSigner();
		const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer);
		const txResult = document.getElementById('tx-result');
		try {
			// setLoading(true);
			const tx = await contract.bid(ethers.toBigInt(quantity), txParams);
			await tx.wait();
			alert('✅ Transaction completed');
		} catch (error) {
			alert("❌ Transaction reverted!");

			console.log(error);
		} finally {
			// setLoading(false);
		}
	}

  const handleConfirm = () => {
    onEntrySubmit();
    // send tx
	buyTickets();
	if (!error) {
      setShowDialog(false);
    }
  };

  return (
    <Card className="w-[1200px] h-[300px] p-8 bg-white dark:bg-gray-800 flex flex-col items-center justify-center space-y-6">
      <div className="text-center space-y-2">
        <h2 className="text-3xl font-bold">Current Pool: {poolAmount} ETH</h2>
        <p className="text-gray-500">Ticket price: {ticketPrice} ETH</p>
        <p className="text-gray-500">Your tickets: {ticketsPerAccount}</p>
      </div>

      <QuantitySelector quantity={quantity} setQuantity={setQuantity}/>

      <Button
        size="lg"
        className="w-48"
        onClick={() => {
          const ethersToSpend = ethers.formatUnits(ethers.toBigInt(quantity) * ethers.parseEther(ticketPrice.split(" ")[0]));
          setEthersToSpend(ethersToSpend);
          setShowDialog(true);      
        }}
        disabled={isProcessing}
      >
        Enter Lottery
      </Button>


      <EntryDialog
        isOpen={showDialog}
        onClose={() => setShowDialog(false)}
        onConfirm={handleConfirm}
        isProcessing={isProcessing}
        error={error}
        ethersToSpend={ethersToSpend}
      />
    </Card>
  );
};

export default EntrySection;
