import { Card } from "@/components/ui/card";
import { Coins, Timer, Trophy } from "lucide-react";

interface FeatureCard {
  title: string;
  description: string;
  icon: React.ReactNode;
  imageUrl: string;
}

const features: FeatureCard[] = [
  {
    title: "Easy Entry",
    description:
      "Enter the lottery with ETH in just a few clicks. No complicated setup required.",
    icon: <Coins className="h-6 w-6 text-purple-500 dark:text-purple-400" />,
    imageUrl:
      "https://images.unsplash.com/photo-1621504450181-5d356f61d307?q=80&w=687&auto=format&fit=crop",
  },
  {
    title: "Fair & Transparent",
    description:
      "Smart contract ensures completely fair and transparent winner selection.",
    icon: <Trophy className="h-6 w-6 text-purple-500 dark:text-purple-400" />,
    imageUrl:
      "https://images.unsplash.com/photo-1639762681485-074b7f938ba0?q=80&w=1032&auto=format&fit=crop",
  },
  {
    title: "Regular Draws",
    description:
      "New lottery rounds every 24 hours. Don't miss your chance to win big!",
    icon: <Timer className="h-6 w-6 text-purple-500 dark:text-purple-400" />,
    imageUrl:
      "https://images.unsplash.com/photo-1639762681485-074b7f938ba0?q=80&w=1032&auto=format&fit=crop",
  },
];

export default function FeatureCards() {
  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
      {features.map((feature, index) => (
        <Card
          key={index}
          className="overflow-hidden hover:shadow-lg transition-all duration-300 hover:-translate-y-1 bg-white dark:bg-gray-800"
        >
          <div className="h-48 overflow-hidden">
            <img
              src={feature.imageUrl}
              alt={feature.title}
              className="w-full h-full object-cover transition-transform duration-300 hover:scale-110"
            />
          </div>
          <div className="p-6">
            <div className="flex items-center gap-3 mb-4">
              {feature.icon}
              <h3 className="text-xl font-semibold">{feature.title}</h3>
            </div>
            <p className="text-muted-foreground">{feature.description}</p>
          </div>
        </Card>
      ))}
    </div>
  );
}
