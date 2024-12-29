import { Github, Twitter, MessageCircle } from "lucide-react";
import { Logo } from "@/components/ui/logo";

export default function Footer() {
  return (
    <footer className="border-t bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container py-8">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          {/* Logo and Description */}
          <div className="space-y-4">
            <Logo />
            <p className="text-sm text-muted-foreground">
              A decentralized lottery platform built on Ethereum. Fair,
              transparent, and secure.
            </p>
          </div>

          {/* Quick Links */}
          <div className="space-y-4">
            <h3 className="font-semibold">Quick Links</h3>
            <ul className="space-y-2">
              <li>
                <a
                  href="#"
                  className="text-sm text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400 transition-colors relative group inline-block"
                >
                  How It Works
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-purple-600 dark:bg-purple-400 transition-all group-hover:w-full"></span>
                </a>
              </li>
              <li>
                <a
                  href="#"
                  className="text-sm text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400 transition-colors relative group inline-block"
                >
                  FAQ
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-purple-600 dark:bg-purple-400 transition-all group-hover:w-full"></span>
                </a>
              </li>
              <li>
                <a
                  href="#"
                  className="text-sm text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400 transition-colors relative group inline-block"
                >
                  Winners
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-purple-600 dark:bg-purple-400 transition-all group-hover:w-full"></span>
                </a>
              </li>
            </ul>
          </div>

          {/* Legal */}
          <div className="space-y-4">
            <h3 className="font-semibold">Legal</h3>
            <ul className="space-y-2">
              <li>
                <a
                  href="#"
                  className="text-sm text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400 transition-colors relative group inline-block"
                >
                  Terms of Service
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-purple-600 dark:bg-purple-400 transition-all group-hover:w-full"></span>
                </a>
              </li>
              <li>
                <a
                  href="#"
                  className="text-sm text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400 transition-colors relative group inline-block"
                >
                  Privacy Policy
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-purple-600 dark:bg-purple-400 transition-all group-hover:w-full"></span>
                </a>
              </li>
              <li>
                <a
                  href="#"
                  className="text-sm text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400 transition-colors relative group inline-block"
                >
                  Smart Contract
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-purple-600 dark:bg-purple-400 transition-all group-hover:w-full"></span>
                </a>
              </li>
            </ul>
          </div>

          {/* Social */}
          <div className="space-y-4">
            <h3 className="font-semibold">Connect</h3>
            <div className="flex space-x-4">
              <a
                href="#"
                className="p-2 rounded-full hover:bg-purple-100 dark:hover:bg-purple-900/50 transition-colors"
              >
                <Twitter className="h-5 w-5 text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400" />
              </a>
              <a
                href="#"
                className="p-2 rounded-full hover:bg-purple-100 dark:hover:bg-purple-900/50 transition-colors"
              >
                <MessageCircle className="h-5 w-5 text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400" />
              </a>
              <a
                href="#"
                className="p-2 rounded-full hover:bg-purple-100 dark:hover:bg-purple-900/50 transition-colors"
              >
                <Github className="h-5 w-5 text-muted-foreground hover:text-purple-600 dark:hover:text-purple-400" />
              </a>
            </div>
          </div>
        </div>

        {/* Copyright */}
        <div className="border-t mt-8 pt-6 flex items-center justify-between">
          <p className="text-sm text-muted-foreground">
            © 2024 ETH Lottery. All rights reserved.
          </p>
          <p className="text-sm text-muted-foreground">
            Built with ❤️ by the ETH Lottery team
          </p>
        </div>
      </div>
    </footer>
  );
}
