import { ThemeProvider } from "@/components/theme/provider";
import { cn } from "@/lib/utils";

import type { Metadata } from "next";
import { Inter } from "next/font/google";

import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Cookie',
  description: 'An awesome Discord Bot.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={cn(inter.className, "min-h-screen w-full bg-neutral-100 dark:bg-neutral-900")}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          {children}
        </ThemeProvider>
      </body>
    </html>
  )
}
