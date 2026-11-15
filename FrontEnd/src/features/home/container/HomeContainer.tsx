import { About } from "../components/AboutSection";
import { Benefits } from "../components/BenefitSection";
import { Cta } from "../components/CTASection";
import { Hero } from "../components/HeroSection";
import { HowTo } from "../components/HowToSection";


export default function Home() {
  return (
    <div className="flex min-h-dvh flex-col bg-background">
      <main className="flex-1">
        <Hero />      
        <HowTo />     
        <Benefits />  
        <About />     
        <Cta />       
      </main>
    </div>
  )
}
