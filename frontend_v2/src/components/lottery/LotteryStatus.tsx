import React from "react";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Timer, Coins } from "lucide-react";

interface LotteryStatusProps {
  poolAmount?: string;
  timeRemaining?: {
    hours: number;
    minutes: number;
    seconds: number;
  };
}

const defaultTimeRemaining = {
  hours: 23,
  minutes: 45,
  seconds: 30,
};

const LotteryStatus = ({
  poolAmount = "10.5 ETH",
  timeRemaining = defaultTimeRemaining,
}: LotteryStatusProps) => {
  return (
    <Card className="w-[1200px] h-[200px] p-6 bg-white dark:bg-gray-800">
      <div className="flex justify-between items-center h-full">
        <div className="flex-1 flex flex-col items-center justify-center border-r border-gray-200 h-full">
          <div className="flex items-center gap-3 mb-4">
            <Coins className="w-8 h-8 text-primary" />
            <h2 className="text-2xl font-bold">All time reward</h2>
          </div>
          <div className="text-4xl font-bold text-primary">{poolAmount} ETH</div>
        </div>

        <div className="flex-1 flex flex-col items-center justify-center h-full">
          <div className="flex items-center gap-3 mb-4">
            <Timer className="w-8 h-8 text-primary" />
            <h2 className="text-2xl font-bold">Time Remaining</h2>
          </div>
          <div className="flex gap-4">
            <div className="text-center">
              <div className="text-4xl font-bold text-primary">
                {timeRemaining.hours}
              </div>
              <Badge variant="secondary" className="mt-2">
                Hours
              </Badge>
            </div>
            <div className="text-4xl font-bold text-primary">:</div>
            <div className="text-center">
              <div className="text-4xl font-bold text-primary">
                {timeRemaining.minutes}
              </div>
              <Badge variant="secondary" className="mt-2">
                Minutes
              </Badge>
            </div>
            <div className="text-4xl font-bold text-primary">:</div>
            <div className="text-center">
              <div className="text-4xl font-bold text-primary">
                {timeRemaining.seconds}
              </div>
              <Badge variant="secondary" className="mt-2">
                Seconds
              </Badge>
            </div>
          </div>
        </div>
      </div>
    </Card>
  );
};

export default LotteryStatus;
