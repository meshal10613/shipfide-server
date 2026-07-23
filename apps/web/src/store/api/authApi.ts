import { apiSlice } from './apiSlice';

export interface RegisterReq {
  name: string;
  email: string;
  phone?: string;
  password: string;
}

export interface LoginReq {
  email: string;
  password: string;
}

export interface VerifyEmailReq {
  email: string;
  otp: string;
}

export interface ForgotPasswordReq {
  email: string;
}

export interface VerifyOtpReq {
  email: string;
  otp: string;
}

export interface ResetPasswordReq {
  email: string;
  otp: string;
  password: string;
}

export interface ChangePasswordReq {
  oldPassword: string;
  newPassword: string;
}

export const authApi = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    register: builder.mutation<{ success: boolean; message: string; data?: any }, RegisterReq>({
      query: (body) => ({
        url: '/auth/register',
        method: 'POST',
        body,
      }),
    }),
    login: builder.mutation<{ success: boolean; message: string; data?: any }, LoginReq>({
      query: (body) => ({
        url: '/auth/login',
        method: 'POST',
        body,
      }),
      invalidatesTags: ['User'],
    }),
    verifyEmail: builder.mutation<{ success: boolean; message: string }, VerifyEmailReq>({
      query: (body) => ({
        url: '/auth/verify-email',
        method: 'POST',
        body,
      }),
    }),
    forgotPassword: builder.mutation<{ success: boolean; message: string }, ForgotPasswordReq>({
      query: (body) => ({
        url: '/auth/forgot-password',
        method: 'POST',
        body,
      }),
    }),
    verifyOtp: builder.mutation<{ success: boolean; message: string }, VerifyOtpReq>({
      query: (body) => ({
        url: '/auth/verify-otp',
        method: 'POST',
        body,
      }),
    }),
    resetPassword: builder.mutation<{ success: boolean; message: string }, ResetPasswordReq>({
      query: (body) => ({
        url: '/auth/reset-password',
        method: 'POST',
        body,
      }),
    }),
    changePassword: builder.mutation<{ success: boolean; message: string }, ChangePasswordReq>({
      query: (body) => ({
        url: '/auth/change-password',
        method: 'POST',
        body,
      }),
    }),
    getMe: builder.query<{ success: boolean; data: any }, void>({
      query: () => '/auth/me',
      providesTags: ['User'],
    }),
    logout: builder.mutation<{ success: boolean; message: string }, void>({
      query: () => ({
        url: '/auth/logout',
        method: 'POST',
      }),
      invalidatesTags: ['User'],
    }),
  }),
});

export const {
  useRegisterMutation,
  useLoginMutation,
  useVerifyEmailMutation,
  useForgotPasswordMutation,
  useVerifyOtpMutation,
  useResetPasswordMutation,
  useChangePasswordMutation,
  useGetMeQuery,
  useLogoutMutation,
} = authApi;
