import { Navbar } from "@/features/home/components/Navbar";
import { Hero } from "@/features/home/components/Hero";
import { HowTo } from "../components/HowTo";
import { Benefits } from "../components/Benefit";
import { About } from "../components/About";
import { Cta } from "../components/CTA";
import { Footer } from "@/shared/components/Footer/Footer";

export default function Home() {
  return (
    <div className="bg-background flex flex-col">
      <Navbar />
      <main className="flex-1">
        <Hero />
        <HowTo />
        <Benefits />
        <About />
        <Cta />
      </main>
      <Footer />
    </div>
  );
}
