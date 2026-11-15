const RADII = [360, 450, 540, 630, 720, 810, 900]

export function FooterCurves() {
  return (
    <div
      aria-hidden
      className="pointer-events-none absolute inset-0 z-0 overflow-hidden"
      style={{
        maskImage:
          "radial-gradient(120% 100% at 78% 118%, #000 0%, rgba(0,0,0,0.55) 42%, transparent 78%)",
        WebkitMaskImage:
          "radial-gradient(120% 100% at 78% 118%, #000 0%, rgba(0,0,0,0.55) 42%, transparent 78%)",
      }}
    >
      <svg
        viewBox="0 0 1440 420"
        preserveAspectRatio="xMidYMax slice"
        className="absolute inset-0 size-full"
        fill="none"
      >
        <g stroke="#ffffff" strokeWidth="1">
          {RADII.map((r, i) => (
            <circle
              key={r}
              cx="1080"
              cy="470"
              r={r}
              strokeOpacity={0.16 - i * 0.014}
            />
          ))}
        </g>

        <g stroke="#e6f3fb" strokeWidth="0.75">
          {RADII.map((r, i) => (
            <circle
              key={r}
              cx="180"
              cy="520"
              r={r * 1.15}
              strokeOpacity={0.09 - i * 0.008}
            />
          ))}
        </g>

        <circle
          cx="1080"
          cy="470"
          r="540"
          stroke="#ffffff"
          strokeOpacity="0.14"
          strokeWidth="2"
          strokeDasharray="1 10"
          strokeLinecap="round"
        />
      </svg>

      {/* glow lembut di belakang busur */}
      <div className="absolute -top-40 right-0 size-[520px] rounded-full bg-[#0d6684]/35 blur-[140px]" />
      <div className="absolute -bottom-32 left-0 size-[420px] rounded-full bg-white/5 blur-[120px]" />
    </div>
  )
}