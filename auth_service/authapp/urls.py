from django.urls import path
from .views import RegisterView, VerifyTokenView, CustomTokenObtainPairView, LogoutView
from rest_framework_simplejwt.views import (TokenObtainPairView, TokenRefreshView)

urlpatterns = [
    path('register/', RegisterView.as_view(), name='register'),
    path('login/', CustomTokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('refresh/', TokenRefreshView.as_view(), name='token_refresh'),
    path('verify-token/', VerifyTokenView.as_view(), name='verify_token'),
    path('logout/', LogoutView.as_view(), name='logout_view'),
]