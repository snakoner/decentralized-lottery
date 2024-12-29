import { Coins } from "lucide-react";

export function Logo() {
  return (
    <div className="flex items-center gap-2 hover:scale-105 transition-transform cursor-pointer">
      <div className="relative">
        <Coins className="w-8 h-8 text-purple-500 dark:text-purple-400" />
        <div className="absolute inset-0 bg-gradient-to-r from-purple-500 via-blue-500 to-purple-500 dark:from-purple-400 dark:via-blue-400 dark:to-purple-400 opacity-50 blur-lg -z-10" />
      </div>
      <h2 className="text-3xl font-bold bg-gradient-to-r from-purple-500 via-blue-500 to-purple-500 dark:from-purple-400 dark:via-blue-400 dark:to-purple-400 bg-clip-text text-transparent animate-gradient-x">
        ETH Lottery
      </h2>
    </div>
  );
}
