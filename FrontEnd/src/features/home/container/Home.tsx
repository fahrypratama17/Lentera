import { Navbar } from "@/features/home/components/Navbar";
import { Hero } from "@/features/home/components/Hero";
import { About } from "../components/AboutSection";
import { Benefits } from "../components/BenefitSection";
import { Cta } from "../components/CTASection";
import { HowTo } from "../components/HowToSection";
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
