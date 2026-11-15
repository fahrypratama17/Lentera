type Ring = {
  tint: string
  opacity: number
  delay: string
  blur: string
}

const RINGS: Ring[] = [
  { tint: "255 255 255", opacity: 0.68, delay: "0s",  blur: "16px" },
  { tint: "236 248 255", opacity: 0.56, delay: "2s",  blur: "15px" },
  { tint: "214 240 255", opacity: 0.46, delay: "4s",  blur: "14px" },
  { tint: "245 252 255", opacity: 0.37, delay: "6s",  blur: "12px" },
  { tint: "255 255 255", opacity: 0.28, delay: "8s",  blur: "11px" },
  { tint: "236 248 255", opacity: 0.20, delay: "10s", blur: "10px" },
]

export function WaterRipple({
  rings = RINGS,
  duration = "12s",
  className = "",
}: {
  rings?: Ring[]
  duration?: string
  className?: string
}) {
  return (
    <div
      aria-hidden
      className={`pointer-events-none absolute inset-0 z-0 grid place-items-center overflow-hidden ${className}`}
    >
      <div
        className="absolute size-[50%] h-[80vh] rounded-full bg-white/30 blur-[40px]"
        style={{ animation: "drop-pulse 6s ease-in-out infinite" }}
      />

      {rings.map((r, i) => (
        <div
          key={i}
          className="absolute aspect-square w-[150%] rounded-full will-change-[transform,opacity]"
          style={{
            background: `radial-gradient(circle at 50% 50%,
              rgba(${r.tint}, 0) 0%,
              rgba(${r.tint}, 0) 66%,
              rgba(${r.tint}, ${r.opacity * 0.18}) 72%,
              rgba(${r.tint}, ${r.opacity * 0.55}) 77%,
              rgba(${r.tint}, ${r.opacity}) 80%,
              rgba(${r.tint}, ${r.opacity * 0.55}) 83%,
              rgba(${r.tint}, ${r.opacity * 0.18}) 88%,
              rgba(${r.tint}, 0) 94%)`,
            filter: `blur(${r.blur})`,
            animation: `water-ripple ${duration} cubic-bezier(.25,.6,.35,1) infinite`,
            animationDelay: r.delay,
            ["--ring-o" as string]: r.opacity,
          }}
        />
      ))}
    </div>
  )
}