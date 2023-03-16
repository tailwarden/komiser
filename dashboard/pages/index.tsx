function Home() {
  const router = useRouter();

  useEffect(() => {
    router.push('/dashboard');
  }, []);
}

export default Home;
