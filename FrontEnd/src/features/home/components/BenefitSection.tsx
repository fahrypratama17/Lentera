import { Diamond } from "lucide-react"
import { Card, CardContent } from "@/shared/components/ui/card"
import { BENEFITS } from "../data/data"

export function Benefits() {
  return (
    <section
      id="manfaat"
      className="bg-linear-to-b from-background via-surface to-background py-[clamp(3.5rem,7vw,6rem)] px-4"
    >
      <div className="container mx-auto">
        <div className="mb-[clamp(2rem,4vw,3.25rem)] flex flex-col items-center gap-3.5 text-center">
          <h2 className="text-[clamp(1.625rem,3.4vw,2.375rem)] leading-tight font-extrabold tracking-[-0.02em] text-balance text-primary-text">
            MANFAAT
          </h2>
        </div>

        <div className="grid grid-cols-1 gap-[18px] md:grid-cols-3">
          {BENEFITS.map(({ icon: Icon, title, text }) => (
            <Card
              key={title}
              className="flex flex-col gap-3.5 rounded-3xl border-border bg-background p-6 pb-7 shadow-none transition hover:-translate-y-1 hover:border-secondary-light hover:shadow-[0_14px_34px_-16px_rgba(10,23,43,0.22)]"
            >
              <span className="grid size-10 place-items-center rounded-xl bg-secondary-light text-primary">
                <Icon className="size-[19px]" strokeWidth={1.9} />
              </span>
              <CardContent className="p-0">
                <h3 className="text-base font-bold tracking-[-0.01em] text-primary-text">{title}</h3>
                <p className="mt-1.5 text-sm leading-relaxed text-secondary-text">{text}</p>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </section>
  )
}
