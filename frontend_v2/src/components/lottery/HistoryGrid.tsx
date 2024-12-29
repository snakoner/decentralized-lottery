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
  winners?: Winner[];
}

const defaultWinners: Winner[] = [
  {
    id: "w1",
    address: "0xabcd...ef12",
    amount: "10 ETH",
    timestamp: "2024-01-19 12:00",
  },
  {
    id: "w2",
    address: "0x7890...cd34",
    amount: "8 ETH",
    timestamp: "2024-01-18 12:00",
  },
  {
    id: "w3",
    address: "0x4567...89ab",
    amount: "12 ETH",
    timestamp: "2024-01-17 12:00",
  },
];

const HistoryGrid = ({ winners = defaultWinners }: HistoryGridProps) => {
  return (
    <Card className="w-[1200px] h-[400px] p-6 bg-gray-50 dark:bg-black">
      <div className="flex justify-between gap-8 h-full">
        <HistoryList />
        <HistoryList
          title="Recent Winners"
          entries={winners.map((winner) => ({
            id: winner.id,
            address: winner.address,
            amount: winner.amount,
            timestamp: winner.timestamp,
          }))}
        />
      </div>
    </Card>
  );
};

export default HistoryGrid;
