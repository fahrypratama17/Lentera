"use client";

import Image from "next/image";
import { useAboutAnimation } from "@/features/home/hooks/useAboutAnimation";

export function About() {
  const { sectionRef } = useAboutAnimation();

  return (
    <section
      ref={sectionRef}
      id="tentang"
      className="bg-background scroll-mt-28 px-4 py-[clamp(4rem,8vw,7rem)]"
    >
      <div className="container mx-auto grid items-center gap-12 lg:grid-cols-2 lg:gap-16">
        <div className="flex flex-col gap-7">
          <div data-about-intro className="flex flex-col items-start gap-5">
            <h2 className="text-primary-text max-w-[22ch] text-[clamp(2rem,4vw,3.25rem)] leading-[1.15] font-extrabold tracking-tight text-balance">
              Setiap orang punya ritme belajar.
              <span className="text-secondary">
                {" "}
                Lentera membantu semuanya tetap searah.
              </span>
            </h2>
          </div>

          <div
            data-about-description
            className="flex max-w-[62ch] flex-col gap-4"
          >
            <p className="text-secondary-text text-base leading-[1.85]">
              Dalam workshop dan kelompok belajar BCC, setiap anggota punya
              perjalanan yang berbeda. Ada yang sudah menyelesaikan materi, ada
              yang baru memulai, dan ada pula yang sedang menunggu review dari
              mentor.
            </p>

            <p className="text-secondary-text text-base leading-[1.85]">
              Di sinilah Lentera hadir. Sebuah ruang belajar bersama yang
              menyatukan materi, tugas, dan progres dalam satu tempat. Mentor
              bisa mendampingi dengan lebih terarah, sementara anggota tahu
              langkah berikutnya yang perlu dilakukan.
            </p>
          </div>
        </div>

        <div
          data-about-visual
          className="bg-secondary-light border-secondary/10 relative isolate order-first flex min-h-95 items-center justify-center overflow-hidden rounded-[2rem] border lg:order-none lg:min-h-145"
        >
          {/* Background pattern */}
          <div
            aria-hidden="true"
            className="pointer-events-none absolute inset-0 opacity-40"
            style={{
              backgroundImage: "radial-gradient(#85B8D0 1px, transparent 1px)",
              backgroundSize: "26px 26px",
            }}
          />

          <div className="absolute top-[12%] left-[12%] size-32 rounded-full bg-white/40 blur-2xl" />
          <div className="bg-secondary/10 absolute right-[8%] bottom-[10%] size-48 rounded-full blur-3xl" />

          <Image
            src="/mascot.webp"
            alt="Maskot burung hantu Lentera"
            width={600}
            height={600}
            sizes="(max-width: 1024px) 80vw, 40vw"
            className="relative z-10 w-[75%] max-w-117.5 object-contain drop-shadow-[0_25px_30px_rgba(31,80,109,0.12)] transition-transform duration-500 hover:scale-105"
          />
        </div>
      </div>
    </section>
  );
}
