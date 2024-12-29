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

interface EntryDialogProps {
  isOpen?: boolean;
  onClose?: () => void;
  onConfirm?: () => void;
  isProcessing?: boolean;
  error?: string;
}

interface EntrySectionProps {
  poolAmount?: string;
  entryFee?: string;
  isProcessing?: boolean;
  error?: string;
  onEntrySubmit?: () => void;
}

const EntryDialog = ({
  isOpen = true,
  onClose = () => {},
  onConfirm = () => {},
  isProcessing = false,
  error = "",
}: EntryDialogProps) => {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Confirm Lottery Entry</DialogTitle>
          <DialogDescription>
            You are about to enter the lottery. This action cannot be undone.
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
  entryFee = "0.1 ETH",
  isProcessing = false,
  error = "",
  onEntrySubmit = () => {},
}: EntrySectionProps) => {
  const [showDialog, setShowDialog] = React.useState(false);

  const handleConfirm = () => {
    onEntrySubmit();
    if (!error) {
      setShowDialog(false);
    }
  };

  return (
    <Card className="w-[1200px] h-[300px] p-8 bg-white dark:bg-gray-800 flex flex-col items-center justify-center space-y-6">
      <div className="text-center space-y-2">
        <h2 className="text-3xl font-bold">Current Pool: {poolAmount}</h2>
        <p className="text-gray-500">Entry Fee: {entryFee}</p>
      </div>

      <Button
        size="lg"
        className="w-48"
        onClick={() => setShowDialog(true)}
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
      />
    </Card>
  );
};

export default EntrySection;
