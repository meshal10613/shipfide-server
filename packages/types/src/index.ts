// Shared TypeScript DTO definitions for Shipfide apps

export type UserRole = 'SUPER_ADMIN' | 'ADMIN' | 'MERCHANT' | 'RIDER';
export type UserStatus = 'PENDING' | 'ACTIVE' | 'SUSPENDED' | 'INACTIVE';

export type ShipmentStatus =
  | 'PENDING'
  | 'PICKUP_RIDER_ASSIGNED'
  | 'PICKED_UP'
  | 'AT_MERCHANT_HUB'
  | 'IN_TRANSIT'
  | 'AT_DELIVERY_HUB'
  | 'OUT_FOR_DELIVERY'
  | 'DELIVERED'
  | 'FAILED_DELIVERY'
  | 'RETURN_INITIATED'
  | 'RETURNED_TO_MERCHANT'
  | 'CANCELLED';

export interface UserResponse {
  id: string;
  name: string;
  email: string;
  phone?: string;
  image?: string;
  role: UserRole;
  status: UserStatus;
  needsPasswordChange: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface AuthResponse {
  accessToken: string;
  refreshToken: string;
  sessionToken: string;
  user: UserResponse;
}
