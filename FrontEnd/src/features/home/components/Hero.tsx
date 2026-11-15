"use client";

import { ArrowRight, BookOpen, Sparkles, TrendingUp } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { HeroWave } from "@/features/home/components/HeroWave";
import Link from "next/link";
import { useHeroReveal } from "@/features/home/hooks/useHeroReveal";
import Image from "next/image";

export function Hero() {
  const { heroRef, leftRef, rightRef, contentRef } = useHeroReveal();

  return (
    <section
      ref={heroRef}
      id="top"
      className="bg-background relative isolate mx-auto min-h-screen overflow-hidden px-4 pt-40 pb-20 md:pt-32"
    >
      <div ref={contentRef} className="flex items-center justify-center">
        <div className="bg-secondary-light relative z-10 container flex min-h-162.5 w-full flex-col items-center justify-center overflow-hidden rounded-[2rem] px-6 py-28 text-center md:px-12">
          <div
            aria-hidden="true"
            className="pointer-events-none absolute inset-0 opacity-40"
            style={{
              backgroundImage: "radial-gradient(#85B8D0 1px, transparent 1px)",
              backgroundSize: "28px 28px",
              maskImage:
                "linear-gradient(to bottom, transparent, black 30%, transparent 90%)",
            }}
          />

          <div className="relative z-20 flex flex-col items-center gap-7">
            <h1 className="text-primary-text max-w-[22ch] text-[clamp(2.5rem,5.5vw,5.5rem)] leading-[1.08] font-extrabold tracking-tight text-balance">
              Belajar bareng,
              <br />
              <span className="text-secondary">berkembang lebih jauh.</span>
            </h1>

            <p className="text-secondary-text max-w-[62ch] text-base leading-relaxed text-pretty md:text-lg">
              Satu ruang untuk berbagi materi, menyelesaikan tugas, dan memantau
              progres belajar. Bikin kolaborasi mentor dan anggota jadi lebih
              mudah, terarah, dan menyenangkan.
            </p>

            <div className="relative z-30 mt-2 grid w-full max-w-90 grid-cols-1 gap-3 sm:max-w-90 sm:grid-cols-2">
              <Button
                size="lg"
                className="group bg-primary shadow-primary/15 hover:bg-secondary h-12 w-full cursor-pointer rounded-full px-5 text-sm font-semibold text-white shadow-lg transition-all duration-300 hover:-translate-y-1 hover:shadow-xl"
              >
                <Link
                  href="/auth"
                  className="flex items-center justify-center gap-2 whitespace-nowrap"
                >
                  Mulai Belajar
                  <ArrowRight className="size-4 shrink-0 transition-transform group-hover:translate-x-1" />
                </Link>
              </Button>

              <Button
                size="lg"
                variant="outline"
                className="border-primary/15 text-primary-text hover:bg-primary h-12 w-full cursor-pointer rounded-full bg-white/90 px-5 text-sm font-semibold transition-colors hover:-translate-y-1 hover:text-white"
              >
                <Link
                  href="#how-to"
                  className="flex items-center justify-center whitespace-nowrap"
                >
                  Jelajahi Fitur
                </Link>
              </Button>
            </div>
          </div>

          <HeroWave />
        </div>
      </div>

      <div
        aria-hidden="true"
        className="pointer-events-none fixed inset-0 z-100"
      >
        <div
          ref={leftRef}
          className="absolute inset-y-0 left-0 w-1/2 overflow-hidden bg-white"
        >
          <div className="absolute inset-y-0 left-0 w-[200%]">
            <Image
              src="/mascot.webp"
              alt="mascot"
              fill
              priority
              sizes="100vw"
              className="object-contain p-[5vw]"
            />
          </div>
        </div>

        <div
          ref={rightRef}
          className="absolute inset-y-0 right-0 w-1/2 overflow-hidden bg-white"
        >
          <div className="absolute inset-y-0 right-0 w-[200%]">
            <Image
              src="/mascot.webp"
              alt="mascot"
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
