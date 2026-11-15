import { Settings } from "lucide-react"
import { Separator } from "@base-ui/react/separator"
import { OwlArt } from "./owl-art"
import { STATS } from "../data/data"
import Image from "next/image"

export function About() {
  return (
    <section id="tentang" className="bg-background py-[clamp(3.5rem,7vw,6rem)] px-4">
      <div className="container grid items-center gap-[clamp(1.75rem,5vw,4rem)] lg:grid-cols-[1.05fr_.95fr] mx-auto">
        <div className="flex flex-col gap-8">

          <h2 className="max-w-[26ch] text-[clamp(1.625rem,3.4vw,2.375rem)] leading-[1.16] font-extrabold tracking-[-0.02em] text-balance text-primary-text">
            Semua Tentang Lentera.
          </h2>

          <p className="text-[15px] leading-[1.72] text-secondary-text text-justify">
            Workshop dan kelompok belajar BCC diikuti banyak anggota dengan ritme yang tidak sama.
            Sebagian sudah menyelesaikan materi pertama, sebagian belum mulai, sebagian menunggu
            keputusan review. Tanpa catatan bersama, informasi ini tersebar di percakapan dan ingatan. Learning Tracker menyatukan materi, progres, dan hasil belajar dalam satu sumber yang sama. Mentor buka satu halaman, langsung tahu siapa yang perlu didorong, siapa yang menunggu review, dan materi mana yang mandek. Tanpa perlu tanya ulang.
          </p>

          <Separator className="mt-2.5 bg-border" />

          <div className="hidden md:grid grid-cols-1 gap-4 pt-1.5 sm:grid-cols-3 sm:gap-3">
            {STATS.map((s) => (
              <div key={s.label}>
                <div className="mb-1 text-[11.5px] font-bold tracking-[0.1em] text-secondary-text/60 uppercase">
                  {s.label}
                </div>
                <div className="text-[22px] font-extrabold tracking-[-0.02em] text-primary">
                  {s.num}
                </div>
                <div className="mt-0.5 text-[13px] leading-snug text-secondary-text">
                  {s.desc}
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="order-first lg:order-none h-full flex items-center justify-center border bg-secondary-light rounded-3xl">
          <Image src={"/owl-art.webp"} alt="maskot" width={500} height={500} className="w-[50%]"/>
        </div>
      </div>
    </section>
  )
}