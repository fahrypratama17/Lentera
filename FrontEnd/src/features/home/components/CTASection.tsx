import { ArrowRight } from "lucide-react"
import { Button } from "@/shared/components/ui/button"
import { WaterRipple } from "./ripple"
import Link from "next/link"

export function Cta() {
  return (
    <section
      id="mulai"
      className="relative overflow-hidden bg-background px-4 pt-16 pb-28 sm:pt-24 sm:pb-32"
    >
      <div className="mx-auto flex w-full items-center justify-center">
        <div className="container relative isolate h-[80vh] flex w-full flex-col items-center justify-center overflow-hidden rounded-3xl bg-secondary  text-center ">
          <WaterRipple />

          <div className="relative z-10 flex flex-col items-center gap-6">
            <h2 className="max-w-[20ch] text-[clamp(2rem,5.2vw,5rem)] leading-tight font-extrabold tracking-[-0.02em] text-balance text-white">
              Mulai dari satu cohort dulu. Anggota menyusul.
            </h2>

            <p className="max-w-[54ch] text-md leading-relaxed text-white/80">
              Cukup nama, email, dan kata sandi. Nggak perlu verifikasi email buat mulai.
            </p>

            <div className="mt-1.5 flex flex-wrap justify-center gap-3">
              <Button
                size="lg"
                className="group h-11 cursor-pointer rounded-full bg-primary px-6 text-[15px] font-semibold text-white hover:bg-white hover:text-primary-text"
              >
                <Link href="/auth" className="flex gap-1 items-center">
                  Mulai sekarang
                  <ArrowRight className="ml-1 size-4 transition-transform group-hover:translate-x-0.5" />
                </Link>
              </Button>

              <Button
                size="lg"
                variant="outline"
                className="h-11 cursor-pointer rounded-full border-white/25 bg-transparent px-6 text-[15px] font-semibold text-white hover:border-white hover:bg-white hover:text-primary-text"
              >
                <Link href="/auth">Daftar</Link>
              </Button>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}