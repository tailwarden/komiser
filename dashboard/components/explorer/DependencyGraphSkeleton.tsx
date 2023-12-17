function DependencyGraphSkeleton() {
  return (
    <>
      <div
        data-testid="loading"
        className="relative flex h-full items-center justify-center bg-dependency-graph bg-[length:40px_40px] align-middle"
      >
        <div>
          <div className="h-3 w-24 rounded-lg bg-cyan-200"></div>
          <div className="mt-2"></div>
          <div className="h-3 w-48 rounded-lg bg-cyan-200"></div>
        </div>
      </div>
    </>
  );
}

export default DependencyGraphSkeleton;
