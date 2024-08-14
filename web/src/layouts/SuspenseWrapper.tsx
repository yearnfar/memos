import { Suspense } from "react";
import { Outlet } from "react-router-dom";
import Loading from "@/pages/Loading";

const SuspenseWrapper = () => {
  return (
    <Suspense fallback={<Loading />}>
      <Outlet />
    </Suspense>
  );
};

export default SuspenseWrapper;
