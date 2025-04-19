import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';
import { Home, Swords, User } from 'lucide-react-native';
import PostScreen from '../modules/post/screens/PostScreen';
import MatchScreen from '../modules/match/screens/MatchScreen';
import ProfileScreen from '../modules/profile/screens/ProfileScreen';
import AnimatedTabBarButton from './components/AnimatedTabBarButton';


const Tab = createBottomTabNavigator();

export default function BottomTabNavigator() {
  return (
    <NavigationContainer>
      <Tab.Navigator
        screenOptions={{
          headerShown: false,
          tabBarActiveTintColor: '#000',
          tabBarLabelStyle: { fontSize: 12, marginBottom: 2 },
          tabBarItemStyle: { paddingTop: 6 },
        }}
      >
        <Tab.Screen
          name="首页"
          component={PostScreen}
          options={{
            tabBarIcon: ({ color, size }) => 
              <Home color={color} size={size} />,
            tabBarButton: (props) => <AnimatedTabBarButton {...props} />,
          }}
        />
        <Tab.Screen
          name="对战"
          component={MatchScreen}
          options={{
            tabBarIcon: ({ color, size }) => <Swords color={color} size={size} />,
            tabBarButton: (props) => <AnimatedTabBarButton {...props} />,
          }}
        />
        <Tab.Screen
          name="我的"
          component={ProfileScreen}
          options={{
            tabBarIcon: ({ color, size }) => <User color={color} size={size} />,
            tabBarButton: (props) => <AnimatedTabBarButton {...props} />,
          }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
}
