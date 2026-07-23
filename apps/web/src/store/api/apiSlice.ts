import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { getAccessTokenCookie } from '@/lib/cookies';

export const apiSlice = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:5000/api/v1',
    credentials: 'include', // Automatically passes access_token, refresh_token, session_token cookies
    prepareHeaders: (headers) => {
      // Read access_token from cookie
      if (typeof window !== 'undefined') {
        const token = getAccessTokenCookie();
        if (token) {
          headers.set('Authorization', `Bearer ${token}`);
        }
      }
      return headers;
    },
  }),
  tagTypes: ['User', 'Shipment', 'Merchant', 'Rider', 'Hub', 'Admin', 'Withdrawal'],
  endpoints: (builder) => ({
    getHealthCheck: builder.query<{ success: boolean; message: string }, void>({
      query: () => '/',
    }),
  }),
});

export const { useGetHealthCheckQuery } = apiSlice;
