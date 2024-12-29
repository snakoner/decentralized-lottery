import React from "react";
import LotteryStatus from "./lottery/LotteryStatus";
import EntrySection from "./lottery/EntrySection";
import HistoryGrid from "./lottery/HistoryGrid";
import FeatureCards from "./lottery/FeatureCards";
import Header from "./layout/Header";
import Footer from "./layout/Footer";

interface HomePageProps {
  poolAmount?: string;
  timeRemaining?: {
    hours: number;
    minutes: number;
    seconds: number;
  };
  entryFee?: string;
  isProcessing?: boolean;
  error?: string;
  onEntrySubmit?: () => void;
}

const HomePage = ({
  poolAmount = "100 ETH",
  timeRemaining = {
    hours: 23,
    minutes: 45,
    seconds: 30,
  },
  entryFee = "0.1 ETH",
  isProcessing = false,
  error = "",
  onEntrySubmit = () => {},
}: HomePageProps) => {
  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-b from-gray-50 to-gray-100 dark:from-gray-950 dark:to-gray-900 transition-colors duration-300">
      <Header />
      <main className="flex-1 py-8 px-4 animate-fade-in">
        <div className="max-w-[1200px] mx-auto space-y-12">
          <div className="text-center space-y-4 mb-8">
            <h1 className="text-4xl font-bold bg-gradient-to-r from-purple-500 via-blue-500 to-purple-500 dark:from-purple-400 dark:via-blue-400 dark:to-purple-400 bg-clip-text text-transparent">
              Welcome to ETH Lottery
            </h1>
            <p className="text-xl text-muted-foreground">
              Your chance to win big with Ethereum
            </p>
          </div>

          <div className="transform hover:scale-[1.01] transition-transform">
            <FeatureCards />
          </div>

          <div className="transform hover:scale-[1.01] transition-transform">
            <LotteryStatus
              poolAmount={poolAmount}
              timeRemaining={timeRemaining}
            />
          </div>

          <div className="transform hover:scale-[1.01] transition-transform">
            <EntrySection
              poolAmount={poolAmount}
              entryFee={entryFee}
              isProcessing={isProcessing}
              error={error}
              onEntrySubmit={onEntrySubmit}
            />
          </div>

          <div className="transform hover:scale-[1.01] transition-transform">
            <HistoryGrid />
          </div>
        </div>
      </main>
      <Footer />
    </div>
  );
};

export default HomePage;
