//https://github.com/ThomasOliver545/angular-material-login-and-register-example/blob/main/src/app/public/interfaces.ts

/*
Interface for the Refresh Token (can look different, based on your backend api)
*/
export interface RefreshToken {
    id: number;
    userId: number;
    token: string;
    refreshCount: number;
    expiryDate: Date;
  }
  
  /*
  Interface for the Login Response (can look different, based on your backend api)
  */
  export interface LoginResponse {
    accessToken: string;
    refreshToken: RefreshToken;
    tokenType: string;
  }
  
  /*
  Interface for the Login Request (can look different, based on your backend api)
  */
  export interface LoginRequest {
    email: string;
    password: string;
  }
  
  /*
  Interface for the Register Request (can look different, based on your backend api)
  */
  export interface RegisterRequest {
    email: string;
    username: string;
    password: string;
  }
  
  /*
  Interface for the Register Response (can look different, based on your backend api)
  */
  export interface RegisterResponse {
    status: number;
    message: string;
  }