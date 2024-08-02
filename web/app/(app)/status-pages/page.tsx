import DataTable from "@/components/status-pages/data-table";
import EmptyState from "@/components/status-pages/empty-state";
import { Button } from "@/components/ui/button";
import { fetchStatusPages } from "@/lib/api/status-pages";
import Link from "next/link";

export default async function StatusPages() {
  const statusPages = await fetchStatusPages();

  return (
    <div className="flex flex-col w-full gap-8">
      <div className="flex items-end justify-between ">
        <div className="flex flex-col gap-1">
          <h1 className="text-2xl font-bold">Status Pages</h1>
          <p className="text-muted-foreground">
            Overview of all your status pages.
          </p>
        </div>

        <Button asChild>
          <Link href="/status-pages/create">Create</Link>
        </Button>
      </div>

      {statusPages?.statusPages.length === 0 ? (
        <EmptyState />
      ) : (
        <DataTable data={statusPages?.statusPages} />
      )}
    </div>
  );
}
