import { getInitials } from "@/features/dashboard/hooks/getInitials";
import { formatDate } from "../utils/format";
import type { CohortMember } from "../types/type";

const MemberRow = ({ member }: { member: CohortMember }) => {
  return (
    <li className="flex items-center gap-3 px-5 py-4">
      <div className="bg-secondary-light text-primary flex size-10 shrink-0 items-center justify-center rounded-full text-sm font-bold">
        {getInitials(member.name)}
      </div>

      <div className="min-w-0 flex-1">
        <p className="text-primary-text flex items-center gap-2 truncate text-sm font-semibold">
          {member.name}
          {member.isMe && (
            <span className="bg-secondary-light text-primary rounded-full px-2 py-0.5 text-[10px] font-bold">
              Kamu
            </span>
          )}
        </p>
        <p className="text-secondary-text truncate text-xs">{member.email}</p>
      </div>

      <p className="text-secondary-text hidden text-xs sm:block">
        Bergabung {formatDate(member.joinedAt)}
      </p>
    </li>
  );
};

const MemberGroup = ({
  title,
  members,
}: {
  title: string;
  members: CohortMember[];
}) => {
  if (members.length === 0) return null;

  return (
    <section className="flex flex-col gap-3">
      <h2 className="text-primary-text text-sm font-bold">{title}</h2>
      <ul className="divide-y divide-[#DCE6ED] overflow-hidden rounded-2xl border border-[#DCE6ED] bg-white">
        {members.map((member) => (
          <MemberRow key={member.id} member={member} />
        ))}
      </ul>
    </section>
  );
};

const MemberList = ({ members }: { members: CohortMember[] }) => {
  const mentors = members.filter((member) => member.role === "mentor");
  const students = members.filter((member) => member.role === "member");

  return (
    <div className="flex flex-col gap-6">
      <MemberGroup title="Mentor" members={mentors} />
      <MemberGroup title={`Anggota (${students.length})`} members={students} />
    </div>
  );
};

export default MemberList;