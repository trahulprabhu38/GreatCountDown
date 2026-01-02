import { useState, useEffect } from 'react';
import { fetchCountdown } from '../services/api';

export const useCountdown = (filter) => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const loadData = async () => {
    try {
      setLoading(true);
      setError(null);
      const result = await fetchCountdown(filter);
      setData(result);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadData();

    // Auto-refresh every 60 seconds
    const interval = setInterval(loadData, 60000);

    return () => clearInterval(interval);
  }, [filter]);

  return { data, loading, error, refetch: loadData };
};
