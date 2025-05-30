from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status, permissions
from .serailizer import RegisterSerializer, CustomTokenObtainSerializer
from rest_framework_simplejwt.tokens import AccessToken, RefreshToken
from rest_framework.exceptions import AuthenticationFailed, ValidationError
from rest_framework_simplejwt.views import TokenObtainPairView


class RegisterView(APIView):
    def post(self, request):
        serializer = RegisterSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response({"message": "User register successfully"}, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class VerifyTokenView(APIView):
    def post(self, request):
        token = request.data.get('token')
        try:
            validated_token = AccessToken(token)
            return Response({
                "valid": True,
                "user_id": validated_token['user_id']
            })
        except Exception as e:
            raise AuthenticationFailed('Invalid token')


class CustomTokenObtainPairView(TokenObtainPairView):
    serializer_class = CustomTokenObtainSerializer


class LogoutView(APIView):
    permission_classes = [permissions.IsAuthenticated]

    def post(self, request):
        try:
            refresh_token = request.data["refresh"],
            token = RefreshToken(refresh_token)
            token.blacklist()
            return Response({
                "message": "Logged out successfully",
                "status": status.HTTP_205_RESET_CONTENT
            })
        except Exception as e:
            raise ValidationError({"error": "Invalid or missing refresh token"})