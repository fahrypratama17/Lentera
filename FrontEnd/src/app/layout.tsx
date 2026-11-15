import type { Metadata } from "next";
import { Nunito } from "next/font/google";
import "../shared/style/globals.css";
import { cn } from "cn";
import Providers from "@/shared/provider/providers";

const nunito = Nunito({
  variable: "--font-nunito",
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "Learning Platform",
  description: "Platform pembelejaran berbasis cohort",
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html
      lang="en"
      className={cn("h-full", "antialiased", nunito.variable, "font-sans")}
    >
      <body className="flex min-h-full flex-col">
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
