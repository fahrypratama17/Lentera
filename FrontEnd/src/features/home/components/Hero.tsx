"use client";

import { ArrowRight } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { WaveBg } from "./wave";
import Link from "next/link";
import { useHeroReveal } from "@/features/home/hooks/useHeroReveal";

export function Hero() {
  const { heroRef, leftRef, rightRef, contentRef } = useHeroReveal();

  return (
    <section
      ref={heroRef}
      id="top"
      className="bg-background relative mx-auto min-h-screen overflow-hidden px-4 pt-40 pb-28 md:pt-32"
    >
      {/* Dark curtain */}
      <div
        aria-hidden="true"
        className="pointer-events-none absolute inset-0 z-30 flex overflow-hidden"
      >
        <div ref={leftRef} className="h-full w-1/2 bg-[#101820]" />
        <div ref={rightRef} className="h-full w-1/2 bg-[#101820]" />
      </div>

      {/* Hero content */}
      <div ref={contentRef} className="flex items-center justify-center">
        <div className="bg-secondary-light relative z-10 container flex w-full flex-col items-center gap-6 overflow-hidden rounded-3xl px-4 py-30 text-center">
          <h1 className="text-primary-text max-w-[30ch] text-[clamp(2.2rem,5.2vw,5rem)] leading-tight font-extrabold tracking-normal text-balance">
            Materi, progres, dan review. Semua di satu tempat.
          </h1>

          <p className="text-secondary-text max-w-[70ch] text-[clamp(0.9375rem,1.5vw,1.0625rem)] leading-relaxed text-pretty">
            Mentor langsung tahu siapa yang belum mulai dan siapa yang menunggu
            review. Anggota tahu persis apa yang dikerjakan berikutnya. Tanpa
            tanya-tanya di grup chat.
          </p>

          <div className="mt-1.5 flex flex-wrap justify-center gap-3">
            <Button
              size="lg"
              className="group bg-primary hover:text-primary-text h-11 cursor-pointer rounded-full px-6 text-[15px] font-semibold text-white hover:bg-white"
            >
              <Link
                href="/auth"
                className="flex items-center justify-center gap-1"
              >
                Mulai
                <ArrowRight className="ml-1 size-4 transition-transform group-hover:translate-x-0.5" />
              </Link>
            </Button>

            <Button
              size="lg"
              variant="outline"
              className="border-border bg-background text-primary-text hover:bg-primary h-11 cursor-pointer rounded-full px-6 text-[15px] font-semibold hover:text-white"
            >
              <Link href="#how-to">Jelajahi</Link>
            </Button>
          </div>

          <WaveBg />
        </div>
      </div>
    </section>
  );
}
