import { ArrowRight } from "lucide-react"
import { Button } from "@/shared/components/ui/button"
import { WaveBg } from "./wave"
import Link from "next/link"

export function Hero() {
  return (
    <section id="top" className="relative overflow-hidden bg-background px-4 pt-14 pb-28 sm:pt-20 sm:pb-32 mx-auto">
      <div className="flex items-center justify-center">
        <div className="w-full container overflow-hidden bg-secondary-light relative z-10 flex flex-col items-center gap-6 text-center px-4 py-30 rounded-3xl">

            <h1 className="max-w-[30ch] text-[clamp(2.2rem,5.2vw,5rem)] leading-tight font-extrabold tracking-normal text-balance text-primary-text">
                Materi, progres, dan review. Semua di satu tempat.
            </h1>

            <p className="max-w-[70ch] text-[clamp(0.9375rem,1.5vw,1.0625rem)] leading-relaxed text-pretty text-secondary-text">
            Mentor langsung tahu siapa yang belum mulai dan siapa yang menunggu review. Anggota tahu persis apa yang dikerjakan berikutnya. Tanpa tanya-tanya di grup chat.
            </p>

            <div className="mt-1.5 flex flex-wrap justify-center gap-3">
            <Button size="lg" className="group h-11 rounded-full px-6 bg-primary text-[15px] font-semibold text-white hover:bg-white hover:text-primary-text cursor-pointer">
                <Link href={"/auth"} className="h-11 flex gap-1 items-center justify-center">
                Mulai
                <ArrowRight className="ml-1 size-4 transition-transform group-hover:translate-x-0.5" />
                </Link>
            </Button>
            <Button
                size="lg"
                variant="outline"
                className="h-11 rounded-full border-border bg-background px-6 text-[15px] font-semibold text-primary-text hover:bg-primary hover:text-white cursor-pointer"
            >
                <Link href={"#how-to"}>
                    Jelajahi
                </Link>
            </Button>
            </div>

            <WaveBg />
        </div>
      </div>
    </section>
  )
}
