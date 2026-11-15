const BannerBackground = () => {
  return (
    <svg
      viewBox="0 0 800 300"
      preserveAspectRatio="none"
      aria-hidden="true"
      className="pointer-events-none absolute inset-0 h-full w-full"
    >
      <circle cx="700" cy="25" r="175" fill="#E6F3FB" opacity=".07" />
      <path
        d="M0 245 C170 135 330 305 520 190 C650 110 715 220 800 120 L800 300 L0 300 Z"
        fill="#79BBD4"
        opacity=".14"
      />
      <path
        d="M0 285 C180 215 340 335 530 255 C650 205 730 275 800 205 L800 300 L0 300 Z"
        fill="#E6F3FB"
        opacity=".12"
      />
    </svg>
  );
};

export default BannerBackground;
