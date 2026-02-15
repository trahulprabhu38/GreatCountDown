const API_BASE_URL = 'http://185.197.251.236:30041/api';

export const fetchCountdown = async (filter = 'year') => {
  let endpoint = `${API_BASE_URL}/countdown`;

  if (filter === 'month') {
    endpoint = `${API_BASE_URL}/countdown/month`;
  } else if (filter === 'week') {
    endpoint = `${API_BASE_URL}/countdown/week`;
  }

  try {
    const response = await fetch(endpoint);
    if (!response.ok) {
      throw new Error('Failed to fetch countdown data');
    }
    return await response.json();
  } catch (error) {
    console.error('API Error:', error);
    throw error;
  }
};
