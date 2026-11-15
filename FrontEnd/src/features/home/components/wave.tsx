export function WaveBg() {
  return (
    <svg
      aria-hidden
      viewBox="0 0 1440 160"
      preserveAspectRatio="none"
      className="pointer-events-none absolute inset-x-0 -bottom-px z-0 h-[90px] w-full sm:h-[130px]"
    >
      <path
        d="M0,96 C220,40 380,140 640,120 C900,100 1080,20 1440,64 L1440,160 L0,160 Z"
        className="fill-secondary"
        opacity={0.85}
      />
      <path
        d="M0,124 C260,74 420,150 700,132 C980,114 1180,56 1440,96 L1440,160 L0,160 Z"
        className="fill-secondary-light"
        opacity={0.55}
      />
    </svg>
  )
}
