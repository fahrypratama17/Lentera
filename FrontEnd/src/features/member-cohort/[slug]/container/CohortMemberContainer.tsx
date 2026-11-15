import { notFound } from "next/navigation";

import MemberList from "../components/MemberList";
import { getCohortMembers } from "../services/cohortDetailService";
import { parseCohortSlug } from "../../utils/slugify";

const CohortMemberContainer = async ({ slug }: { slug: string }) => {
  const key = parseCohortSlug(slug);
  if (!key) notFound();
  const members = await getCohortMembers(key);

  return <MemberList members={members} />;
};

export default CohortMemberContainer;