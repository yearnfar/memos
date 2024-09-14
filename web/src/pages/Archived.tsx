import { Button, Tooltip } from "@mui/joy";
import dayjs from "dayjs";
import { ArchiveIcon, ArchiveRestoreIcon, ArrowDownIcon, TrashIcon } from "lucide-react";
import { ClientError } from "nice-grpc-web";
import { useEffect, useState } from "react";
import toast from "react-hot-toast";
import Empty from "@/components/Empty";
import MemoContent from "@/components/MemoContent";
import MemoFilters from "@/components/MemoFilters";
import MobileHeader from "@/components/MobileHeader";
import SearchBar from "@/components/SearchBar";
import { DEFAULT_LIST_MEMOS_PAGE_SIZE } from "@/helpers/consts";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useMemoFilterStore, useMemoList, useMemoStore } from "@/store/v1";
import { RowStatus } from "@/types/proto/api/v1/common";
import { Memo } from "@/types/proto/api/v1/memo_service";
import { useTranslate } from "@/utils/i18n";

const Archived = () => {
  const t = useTranslate();
  const user = useCurrentUser();
  const memoStore = useMemoStore();
  const memoList = useMemoList();
  const memoFilterStore = useMemoFilterStore();
  const [isRequesting, setIsRequesting] = useState(true);
  const [nextPageToken, setNextPageToken] = useState<string>("");
  const sortedMemos = memoList.value
    .filter((memo) => memo.rowStatus === RowStatus.ARCHIVED)
    .sort((a, b) =>
      memoFilterStore.orderByTimeAsc
        ? dayjs(a.displayTime).unix() - dayjs(b.displayTime).unix()
        : dayjs(b.displayTime).unix() - dayjs(a.displayTime).unix(),
    );

  useEffect(() => {
    memoList.reset();
    fetchMemos("");
  }, [memoFilterStore.filters]);

  const fetchMemos = async (nextPageToken: string) => {
    setIsRequesting(true);
    const filters = [`creator == "${user.name}"`, `row_status == "ARCHIVED"`];
    const contentSearch: string[] = [];
    const tagSearch: string[] = [];
    for (const filter of memoFilterStore.filters) {
      if (filter.factor === "contentSearch") {
        contentSearch.push(`"${filter.value}"`);
      } else if (filter.factor === "tagSearch") {
        tagSearch.push(`"${filter.value}"`);
      }
    }
    if (memoFilterStore.orderByTimeAsc) {
      filters.push(`order_by_time_asc == true`);
    }
    if (contentSearch.length > 0) {
      filters.push(`content_search == [${contentSearch.join(", ")}]`);
    }
    if (tagSearch.length > 0) {
      filters.push(`tag_search == [${tagSearch.join(", ")}]`);
    }
    const response = await memoStore.fetchMemos({
      pageSize: DEFAULT_LIST_MEMOS_PAGE_SIZE,
      filter: filters.join(" && "),
      pageToken: nextPageToken,
    });
    setIsRequesting(false);
    setNextPageToken(response.nextPageToken);
  };

  const handleDeleteMemoClick = async (memo: Memo) => {
    const confirmed = window.confirm(t("memo.delete-confirm"));
    if (confirmed) {
      await memoStore.deleteMemo(memo.name);
    }
  };

  const handleRestoreMemoClick = async (memo: Memo) => {
    try {
      await memoStore.updateMemo(
        {
          name: memo.name,
          rowStatus: RowStatus.ACTIVE,
        },
        ["row_status"],
      );
      toast(t("message.restored-successfully"));
    } catch (error: unknown) {
      console.error(error);
      toast.error((error as ClientError).details);
    }
  };

  return (
    <section className="@container w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      <MobileHeader />
      <div className="w-full px-4 sm:px-6">
        <div className="w-full flex flex-col justify-start items-start">
          <div className="w-full flex flex-row justify-between items-center mb-2">
            <div className="flex flex-row justify-start items-center gap-1">
              <ArchiveIcon className="w-5 h-auto opacity-70 shrink-0" />
              <span>{t("common.archived")}</span>
            </div>
            <div className="w-44">
              <SearchBar />
            </div>
          </div>
          <MemoFilters />
          {sortedMemos.map((memo) => (
            <div
              key={memo.name}
              className="relative flex flex-col justify-start items-start w-full p-4 pt-3 mb-2 bg-white dark:bg-zinc-800 rounded-lg"
            >
              <div className="w-full mb-1 flex flex-row justify-between items-center">
                <div className="w-full max-w-[calc(100%-20px)] flex flex-row justify-start items-center mr-1">
                  <div className="text-sm leading-6 text-gray-400 select-none">
                    <relative-time datetime={memo.displayTime?.toISOString()} tense="past"></relative-time>
                  </div>
                </div>
                <div className="flex flex-row justify-end items-center gap-x-2">
                  <Tooltip title={t("common.restore")} placement="top">
                    <button onClick={() => handleRestoreMemoClick(memo)}>
                      <ArchiveRestoreIcon className="w-4 h-auto cursor-pointer text-gray-500 dark:text-gray-400" />
                    </button>
                  </Tooltip>
                  <Tooltip title={t("common.delete")} placement="top">
                    <button onClick={() => handleDeleteMemoClick(memo)} className="text-gray-500 dark:text-gray-400">
                      <TrashIcon className="w-4 h-auto cursor-pointer" />
                    </button>
                  </Tooltip>
                </div>
              </div>
              <MemoContent key={`${memo.name}-${memo.displayTime}`} memoName={memo.name} nodes={memo.nodes} readonly={true} />
            </div>
          ))}
          {nextPageToken && (
            <div className="w-full flex flex-row justify-center items-center my-4">
              <Button
                variant="plain"
                color="neutral"
                loading={isRequesting}
                endDecorator={<ArrowDownIcon className="w-4 h-auto" />}
                onClick={() => fetchMemos(nextPageToken)}
              >
                {t("memo.load-more")}
              </Button>
            </div>
          )}
          {!nextPageToken && sortedMemos.length === 0 && (
            <div className="w-full mt-12 mb-8 flex flex-col justify-center items-center italic">
              <Empty />
              <p className="mt-2 text-gray-600 dark:text-gray-400">{t("message.no-data")}</p>
            </div>
          )}
        </div>
      </div>
    </section>
  );
};

export default Archived;
