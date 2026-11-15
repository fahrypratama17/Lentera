import { Settings } from "lucide-react";
import { Separator } from "@base-ui/react/separator";
import { OwlArt } from "./owl-art";
import { STATS } from "../data/data";
import Image from "next/image";

export function About() {
  return (
    <section
      id="tentang"
      className="bg-background px-4 py-[clamp(3.5rem,7vw,6rem)]"
    >
      <div className="container mx-auto grid items-center gap-[clamp(1.75rem,5vw,4rem)] lg:grid-cols-[1.05fr_.95fr]">
        <div className="flex flex-col gap-8">
          <h2 className="text-primary-text max-w-[26ch] text-[clamp(1.625rem,3.4vw,2.375rem)] leading-[1.16] font-extrabold tracking-[-0.02em] text-balance">
            Semua Tentang Lentera.
          </h2>

          <p className="text-secondary-text text-justify text-[15px] leading-[1.72]">
            Workshop dan kelompok belajar BCC diikuti banyak anggota dengan
            ritme yang tidak sama. Sebagian sudah menyelesaikan materi pertama,
            sebagian belum mulai, sebagian menunggu keputusan review. Tanpa
            catatan bersama, informasi ini tersebar di percakapan dan ingatan.
            Learning Tracker menyatukan materi, progres, dan hasil belajar dalam
            satu sumber yang sama. Mentor buka satu halaman, langsung tahu siapa
            yang perlu didorong, siapa yang menunggu review, dan materi mana
            yang mandek. Tanpa perlu tanya ulang.
          </p>

          <Separator className="bg-border mt-2.5" />

          <div className="hidden grid-cols-1 gap-4 pt-1.5 sm:grid-cols-3 sm:gap-3 md:grid">
            {STATS.map((s) => (
              <div key={s.label}>
                <div className="text-secondary-text/60 mb-1 text-[11.5px] font-bold tracking-[0.1em] uppercase">
                  {s.label}
                </div>
                <div className="text-primary text-[22px] font-extrabold tracking-[-0.02em]">
                  {s.num}
                </div>
                <div className="text-secondary-text mt-0.5 text-[13px] leading-snug">
                  {s.desc}
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="bg-secondary-light order-first flex h-full items-center justify-center rounded-3xl border lg:order-none">
          <Image
            src={"/mascot.webp"}
            alt="mascot"
            width={500}
            height={500}
            className="w-[50%]"
          />
        </div>
      </div>
    </section>
  );
}
