import React from "react";
import { Card } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

interface HistoryEntry {
  id: string;
  address: string;
  amount: string;
  timestamp: string;
}

interface HistoryListProps {
  title?: string;
  entries?: HistoryEntry[];
}

const defaultEntries: HistoryEntry[] = [
  {
    id: "1",
    address: "0x1234...5678",
    amount: "0.1 ETH",
    timestamp: "2024-01-20 14:30",
  },
  {
    id: "2",
    address: "0x8765...4321",
    amount: "0.2 ETH",
    timestamp: "2024-01-20 14:25",
  },
  {
    id: "3",
    address: "0x9876...1234",
    amount: "0.15 ETH",
    timestamp: "2024-01-20 14:20",
  },
];

const getAvatarUrl = (address: string) => {
  return `https://api.dicebear.com/7.x/identicon/svg?seed=${address}`;
};

const getInitials = (address: string) => {
  return address.slice(2, 4).toUpperCase();
};

const HistoryList = ({
  title = "Entry History",
  entries = defaultEntries,
}: HistoryListProps) => {
  return (
    <Card className="w-[580px] h-[400px] p-4 bg-white dark:bg-gray-800">
      <h3 className="text-xl font-semibold mb-4">{title}</h3>
      <ScrollArea className="h-[320px] w-full">
        <div className="space-y-4">
          {entries.map((entry) => (
            <div
              key={entry.id}
              className="p-3 border rounded-lg flex items-center justify-between hover:bg-gray-50 dark:hover:bg-gray-700/50 group"
            >
              <div className="flex items-center gap-3">
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <Avatar className="h-10 w-10 border-2 border-transparent group-hover:border-purple-500 dark:group-hover:border-purple-400 transition-colors">
                        <img
                          src={getAvatarUrl(entry.address)}
                          alt={entry.address}
                          className="object-cover"
                        />
                        <AvatarFallback>
                          {getInitials(entry.address)}
                        </AvatarFallback>
                      </Avatar>
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>Address: {entry.address}</p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
                <div className="flex flex-col">
                  <span className="text-sm font-medium">{entry.address}</span>
                  <span className="text-xs text-muted-foreground">
                    {entry.timestamp}
                  </span>
                </div>
              </div>
              <Badge
                variant="secondary"
                className="ml-2 group-hover:bg-purple-100 group-hover:text-purple-600 dark:group-hover:bg-purple-900 dark:group-hover:text-purple-400 transition-colors"
              >
                {entry.amount}
              </Badge>
            </div>
          ))}
        </div>
      </ScrollArea>
    </Card>
  );
};

export default HistoryList;
