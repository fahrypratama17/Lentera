"use client";

import { ArrowRight, Sparkles } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { WaterRipple } from "./ripple";
import Link from "next/link";
import { useCtaAnimation } from "@/features/home/hooks/useCtaAnimation";

export function Cta() {
  const { sectionRef } = useCtaAnimation();

  return (
    <section
      ref={sectionRef}
      id="mulai"
      className="bg-background relative scroll-mt-28 overflow-hidden px-4 pt-16 pb-28 sm:pt-24 sm:pb-32"
    >
      <div className="container mx-auto">
        <div className="from-primary via-secondary relative isolate flex min-h-[80svh] w-full flex-col items-center justify-center overflow-hidden rounded-[2rem] bg-linear-to-br to-[#07506F] px-6 py-24 text-center sm:px-12">
          <div
            aria-hidden="true"
            className="pointer-events-none absolute inset-0 z-0"
          >
            <WaterRipple />
          </div>
          <div
            aria-hidden="true"
            className="pointer-events-none absolute -top-40 -right-20 size-100 rounded-full bg-white/10 blur-[100px]"
          />
          <div
            aria-hidden="true"
            className="pointer-events-none absolute -bottom-40 -left-20 size-100 rounded-full bg-[#79C9E8]/20 blur-[100px]"
          />
          <div className="relative z-10 flex w-full max-w-4xl flex-col items-center gap-7">
            <h2
              data-cta-reveal
              className="max-w-[20ch] text-[clamp(2.5rem,5.2vw,5rem)] leading-[1.1] font-extrabold tracking-tight text-balance text-white"
            >
              Siap bikin proses belajar
              <span className="text-[#BDE8FF]"> lebih terarah?</span>
            </h2>
            <p
              data-cta-reveal
              className="max-w-[58ch] text-base leading-relaxed text-pretty text-white/80 md:text-lg"
            >
              Mulai dari satu kelas, lalu kelola materi, tugas, dan progres
              belajar bersama. Semua lebih praktis, tanpa harus berpindah-pindah
              aplikasi.
            </p>
            <div
              data-cta-reveal
              className="mt-3 grid w-full max-w-105 grid-cols-1 gap-3 sm:grid-cols-2"
            >
              <Button
                size="lg"
                className="group text-primary hover:bg-secondary-light h-12 w-full cursor-pointer rounded-full bg-white px-5 text-base font-semibold shadow-lg transition-all duration-300 hover:-translate-y-1 hover:shadow-xl"
              >
                <Link
                  href="/auth"
                  className="flex items-center justify-center gap-2 whitespace-nowrap"
                >
                  Mulai Gratis
                  <ArrowRight className="size-4 shrink-0 transition-transform group-hover:translate-x-1" />
                </Link>
              </Button>

              <Button
                size="lg"
                variant="outline"
                className="h-12 w-full cursor-pointer rounded-full border-white/35 bg-white/10 px-5 text-base font-semibold text-white backdrop-blur-sm transition-all duration-300 hover:-translate-y-1 hover:border-white/70 hover:bg-white/20 hover:text-white"
              >
                <Link
                  href="#cara-menggunakan"
                  className="flex items-center justify-center whitespace-nowrap"
                >
                  Lihat Cara Kerja
                </Link>
              </Button>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
