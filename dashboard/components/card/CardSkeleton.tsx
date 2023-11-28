function CardSkeleton() {
  return (
    <div className="flex h-[7.5rem] animate-pulse items-center rounded-lg bg-white px-6 text-sm">
      <div className="flex w-full gap-6">
        <div className="h-10 w-10 flex-shrink-0 rounded-xl bg-cyan-200"></div>
        <div className="flex w-full flex-col gap-3">
          <div className="h-4 w-[36%] rounded-lg bg-cyan-200"></div>
          <div className="h-4 w-[86%] rounded-lg bg-cyan-200"></div>
        </div>
      </div>
    </div>
  );
}

export default CardSkeleton;
