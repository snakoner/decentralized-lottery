import React from "react";
import HistoryList from "./HistoryList";
import { Card } from "@/components/ui/card";

interface Winner {
  id: string;
  address: string;
  amount: string;
  timestamp: string;
}

interface HistoryGridProps {
  participants?: Winner[];
  winners?: Winner[];
}

const defaultWinners: Winner[] = [

];

const truncateAddress = (address: string) => `${address.slice(0, 6)}...${address.slice(-4)}`.toLowerCase();

const HistoryGrid = ({ participants = defaultWinners, winners = defaultWinners }: HistoryGridProps) => {
  return (
    <Card className="w-[1200px] h-[400px] p-6 bg-gray-50 dark:bg-black">
      <div className="flex justify-between gap-8 h-full">
        <HistoryList 
          entries={participants.map((participant) => ({
            id: participant.id,
            address: truncateAddress(participant.address),
            amount: participant.amount + " ETH",
            timestamp: participant.timestamp,
          }))}        
        />
        <HistoryList
          title="Recent Winners"
          entries={winners.map((winner) => ({
            id: winner.id,
            address: truncateAddress(winner.address),
            amount: winner.amount + " ETH",
            timestamp: winner.timestamp,
          }))}
        />
      </div>
    </Card>
  );
};

export default HistoryGrid;
