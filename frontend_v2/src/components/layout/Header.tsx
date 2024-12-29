import { Button } from "@/components/ui/button";
import { Wallet } from "lucide-react";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Logo } from "@/components/ui/logo";

export default function Header() {
  return (
    <header className="sticky top-0 z-50 border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-16 items-center justify-between">
        <Logo />
        <nav className="flex items-center space-x-4">
          <ThemeToggle />
          <Button
            variant="outline"
            className="gap-2 hover:scale-105 transition-all hover:bg-purple-50 hover:border-purple-200 hover:text-purple-600 dark:hover:bg-purple-900 dark:hover:border-purple-700 dark:hover:text-purple-400"
          >
            <Wallet className="h-4 w-4" />
            Connect Wallet
          </Button>
        </nav>
      </div>
    </header>
  );
}
