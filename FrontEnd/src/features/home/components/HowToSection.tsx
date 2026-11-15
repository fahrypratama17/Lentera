import { Sparkles } from "lucide-react"
import { Card, CardContent } from "@/shared/components/ui/card"
import { STEPS } from "../data/data"
import { cn } from "@/lib/utils"

const VARIANTS = {
  tint:   "bg-secondary-light border-secondary-light",
  plain:  "bg-background border-border",
  active: "bg-primary border-primary text-primary-foreground shadow-[0_18px_40px_-18px_rgba(31,80,109,0.55)]",
} as const

export function HowTo() {
  return (
    <section id="how-to" className="scroll-mt-28 bg-background py-[clamp(3.5rem,7vw,6rem)] px-4">
      <div className="container mx-auto">
        <div className="mb-[clamp(2rem,4vw,3.25rem)] flex flex-col items-center gap-3.5 text-center">
          <h2 className="max-w-[30ch] text-[clamp(1.625rem,3.4vw,2.375rem)] leading-tight font-extrabold tracking-[-0.02em] text-balance text-primary-text">
            Dari cohort kosong sampai materi tuntas, dalam empat langkah.
          </h2>
          <p className="max-w-[56ch] text-sm leading-relaxed text-secondary-text">
            Alurnya sama di semua cohort. Yang beda cuma peranmu di dalamnya.
          </p>
        </div>

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {STEPS.map(({ n, icon: Icon, title, text, variant }) => {
            return (
              <Card
                key={n}
                className={cn(
                  "relative flex flex-col gap-3.5 rounded-3xl border p-6 pb-7 shadow-none transition-transform duration-200 hover:-translate-y-1 hover:shadow-[0_14px_34px_-16px_rgba(10,23,43,0.22)]",
                  VARIANTS[variant]
                )}
              >
                <span className={cn(
                  "absolute top-4 right-5 text-xs font-extrabold tracking-[0.06em] text-secondary"
                )}>
                  {String(n).padStart(2, "0")}
                </span>

                <span className={cn(
                  "grid size-[42px] place-items-center rounded-xl bg-primary-foreground/15 text-secondary"
                )}>
                  <Icon className="size-5" strokeWidth={1.9} />
                </span>

                <CardContent className="p-0">
                  <h3 className={cn(
                    "text-base font-bold tracking-[-0.01em] text-secondary"
                  )}>
                    {title}
                  </h3>
                  <p className={cn(
                    "mt-1.5 text-sm leading-relaxed text-secondary"
                  )}>
                    {text}
                  </p>
                </CardContent>
              </Card>
            )
          })}
        </div>
      </div>
    </section>
  )
}