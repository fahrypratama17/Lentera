import { progressStatusConfig } from "../types/type";
import type { ProgressStatus } from "../types/type";

const ProgressStatusBadge = ({ status }: { status: ProgressStatus }) => {
  const { label, className } = progressStatusConfig[status];

  return (
    <span
      className={`shrink-0 rounded-full px-3 py-1 text-xs font-semibold ${className}`}
    >
      {label}
    </span>
  );
};

export default ProgressStatusBadge;