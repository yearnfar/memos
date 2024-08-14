import { uniq } from "lodash-es";
import { memo, useEffect, useState } from "react";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useUserStore } from "@/store/v1";
import { Memo } from "@/types/proto/api/v1/memo_service";
import { Reaction, Reaction_Type } from "@/types/proto/api/v1/reaction_service";
import { User } from "@/types/proto/api/v1/user_service";
import ReactionSelector from "./ReactionSelector";
import ReactionView from "./ReactionView";

interface Props {
  memo: Memo;
  reactions: Reaction[];
}

const MemoReactionListView = (props: Props) => {
  const { memo, reactions } = props;
  const currentUser = useCurrentUser();
  const userStore = useUserStore();
  const [reactionGroup, setReactionGroup] = useState<Map<Reaction_Type, User[]>>(new Map());

  useEffect(() => {
    (async () => {
      const reactionGroup = new Map<Reaction_Type, User[]>();
      for (const reaction of reactions) {
        const user = await userStore.getOrFetchUserByName(reaction.creator);
        const users = reactionGroup.get(reaction.reactionType) || [];
        users.push(user);
        reactionGroup.set(reaction.reactionType, uniq(users));
      }
      setReactionGroup(reactionGroup);
    })();
  }, [reactions]);

  return (
    reactions.length > 0 && (
      <div className="w-full flex flex-row justify-start items-start flex-wrap gap-1 select-none">
        {Array.from(reactionGroup).map(([reactionType, users]) => {
          return <ReactionView key={`${reactionType.toString()} ${users.length}`} memo={memo} reactionType={reactionType} users={users} />;
        })}
        {currentUser && <ReactionSelector memo={memo} />}
      </div>
    )
  );
};

export default memo(MemoReactionListView);
