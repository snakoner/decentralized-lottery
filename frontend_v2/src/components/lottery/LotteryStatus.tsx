import React from "react";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Timer, Coins, Tornado } from "lucide-react";

interface TimeRemaining {
    hours: string;
    minutes: string;
    seconds: string;
}

interface LotteryStatusProps {
  currentRound?: number;
  poolAmount?: string;
  secondsTimeRemaining?: number;
}

const formatTime = (seconds: number): TimeRemaining => {
  const convertTimeToString = (t: number) => { return t < 10 ? `0${t.toString()}` : t.toString(); }

  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const secs = seconds % 60;

  return {
    hours: convertTimeToString(hours),
    minutes: convertTimeToString(minutes),
    seconds: convertTimeToString(secs),
  };
};

const LotteryStatus = ({
  currentRound = 0,
  poolAmount = "10.5 ETH",
  secondsTimeRemaining = 3600,
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

        <div className="flex-1 flex flex-col items-center justify-center border-r border-gray-200 h-full">
          <div className="flex items-center gap-3 mb-4">
            <Tornado className="w-8 h-8 text-primary" />
            <h2 className="text-2xl font-bold">Round</h2>
          </div>
          <div className="text-4xl font-bold text-primary">{currentRound}</div>
        </div>


        <div className="flex-1 flex flex-col items-center justify-center h-full">
          <div className="flex items-center gap-3 mb-4">
            <Timer className="w-8 h-8 text-primary" />
            <h2 className="text-2xl font-bold">Time Remaining</h2>
          </div>
          <div className="flex gap-4">
            <div className="text-center">
              <div className="text-4xl font-bold text-primary">
                {formatTime(secondsTimeRemaining).hours}
              </div>
              <Badge variant="secondary" className="mt-2">
                Hours
              </Badge>
            </div>
            <div className="text-4xl font-bold text-primary">:</div>
            <div className="text-center">
              <div className="text-4xl font-bold text-primary">
                {formatTime(secondsTimeRemaining).minutes}
              </div>
              <Badge variant="secondary" className="mt-2">
                Minutes
              </Badge>
            </div>
            <div className="text-4xl font-bold text-primary">:</div>
            <div className="text-center">
              <div className="text-4xl font-bold text-primary">
                {formatTime(secondsTimeRemaining).seconds}
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
