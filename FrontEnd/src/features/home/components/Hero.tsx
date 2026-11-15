"use client";

import { ArrowRight } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { WaveBg } from "./wave";
import Link from "next/link";
import { useHeroReveal } from "@/features/home/hooks/useHeroReveal";
import Image from "next/image";

export function Hero() {
  const { heroRef, leftRef, rightRef, contentRef } = useHeroReveal();

  return (
    <section
      ref={heroRef}
      id="top"
      className="bg-background relative isolate mx-auto min-h-screen overflow-hidden px-4 pt-40 pb-28 md:pt-32"
    >
      {/* Hero Content */}
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
              className="group bg-primary hover:text-primary-text h-11 rounded-full px-6 text-[15px] font-semibold text-white hover:bg-white"
            >
              <Link href="/auth">
                Mulai
                <ArrowRight className="ml-1 size-4 transition-transform group-hover:translate-x-0.5" />
              </Link>
            </Button>

            <Button
              size="lg"
              variant="outline"
              className="border-border bg-background text-primary-text hover:bg-primary h-11 rounded-full px-6 text-[15px] font-semibold hover:text-white"
            >
              <Link href="#how-to">Jelajahi</Link>
            </Button>
          </div>

          <WaveBg />
        </div>
      </div>

      {/* Fullscreen split curtain */}
      <div
        aria-hidden="true"
        className="pointer-events-none absolute inset-0 z-50"
      >
        {/* Left half */}
        <div
          ref={leftRef}
          className="absolute inset-y-0 left-0 w-1/2 overflow-hidden bg-white"
        >
          <div className="absolute inset-y-0 left-0 w-[200%]">
            <Image
              src="/mascot.webp"
              alt=""
              fill
              priority
              sizes="100vw"
              className="object-contain p-[5vw]"
            />
          </div>
        </div>

        {/* Right half */}
        <div
          ref={rightRef}
          className="absolute inset-y-0 right-0 w-1/2 overflow-hidden bg-white"
        >
          <div className="absolute inset-y-0 right-0 w-[200%]">
            <Image
              src="/mascot.webp"
              alt=""
              fill
              priority
              sizes="100vw"
              className="object-contain p-[5vw]"
            />
          </div>
        </div>
      </div>
    </section>
  );
}
