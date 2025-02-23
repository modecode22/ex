### app\(auth)\login.tsx

```typescript
import React, { useState } from 'react';
import { View, Text, ScrollView, KeyboardAvoidingView, Platform } from 'react-native';
import { Link } from 'expo-router';
import { FontAwesome } from '@expo/vector-icons';
import { Input } from '../../components/ui/Input';
import { Button } from '../../components/ui/Button';
import { useAuth } from '../../context/AuthContext';
import { colors } from '../../constants/theme';

export default function LoginScreen() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { signIn, isLoading } = useAuth();

  const handleLogin = async () => {
    try {
      setError('');
      await signIn(email, password);
    } catch (err) {
      setError('Invalid email or password');
    }
  };

  return (
    <KeyboardAvoidingView 
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1 bg-gray-50"
    >
      <ScrollView 
        contentContainerStyle={{ flexGrow: 1 }}
        keyboardShouldPersistTaps="handled"
      >
        <View className="flex-1 justify-center px-6 py-12">
          <View className="mb-8 items-center">
            <View className="w-16 h-16 bg-primary-100 rounded-full items-center justify-center mb-4">
              <FontAwesome name="tasks" size={32} color={colors.primary[500]} />
            </View>
            <Text className="text-2xl font-bold text-gray-900 mb-2">Welcome back!</Text>
            <Text className="text-gray-500 text-center">
              Sign in to your account to continue
            </Text>
          </View>

          {error ? (
            <View className="mb-4 p-4 bg-red-50 rounded-lg">
              <Text className="text-red-500 text-center">{error}</Text>
            </View>
          ) : null}

          <View className="space-y-4">
            <Input
              label="Email"
              placeholder="Enter your email"
              value={email}
              onChangeText={setEmail}
              autoCapitalize="none"
              keyboardType="email-address"
            />

            <Input
              label="Password"
              placeholder="Enter your password"
              value={password}
              onChangeText={setPassword}
              secureTextEntry
            />

            <Button
              title="Sign In"
              onPress={handleLogin}
              isLoading={isLoading}
              fullWidth
            />
          </View>

          <View className="mt-8 items-center">
            <Text className="text-gray-500">Don't have an account?</Text>
            <Link href="/(auth)\register" className="mt-2">
              <Text className="text-primary-600 font-medium">Create Account</Text>
            </Link>
          </View>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}

```

### app\(auth)\register.tsx

```typescript
import React, { useState } from 'react';
import { View, Text, ScrollView, KeyboardAvoidingView, Platform } from 'react-native';
import { Link } from 'expo-router';
import { FontAwesome } from '@expo/vector-icons';
import { Input } from '../../components/ui/Input';
import { Button } from '../../components/ui/Button';
import { useAuth } from '../../context/AuthContext';
import { colors } from '../../constants/theme';

export default function RegisterScreen() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState('');
  const { signUp, isLoading } = useAuth();

  const handleRegister = async () => {
    try {
      setError('');
      
      if (password !== confirmPassword) {
        setError('Passwords do not match');
        return;
      }

      await signUp(email, password);
    } catch (err) {
      setError('Failed to create account');
    }
  };

  return (
    <KeyboardAvoidingView 
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1 bg-gray-50"
    >
      <ScrollView 
        contentContainerStyle={{ flexGrow: 1 }}
        keyboardShouldPersistTaps="handled"
      >
        <View className="flex-1 justify-center px-6 py-12">
          <View className="mb-8 items-center">
            <View className="w-16 h-16 bg-primary-100 rounded-full items-center justify-center mb-4">
              <FontAwesome name="user-plus" size={32} color={colors.primary[500]} />
            </View>
            <Text className="text-2xl font-bold text-gray-900 mb-2">Create Account</Text>
            <Text className="text-gray-500 text-center">
              Sign up to get started with Task Manager
            </Text>
          </View>

          {error ? (
            <View className="mb-4 p-4 bg-red-50 rounded-lg">
              <Text className="text-red-500 text-center">{error}</Text>
            </View>
          ) : null}

          <View className="space-y-4">
            <Input
              label="Email"
              placeholder="Enter your email"
              value={email}
              onChangeText={setEmail}
              autoCapitalize="none"
              keyboardType="email-address"
            />

            <Input
              label="Password"
              placeholder="Create a password"
              value={password}
              onChangeText={setPassword}
              secureTextEntry
            />

            <Input
              label="Confirm Password"
              placeholder="Confirm your password"
              value={confirmPassword}
              onChangeText={setConfirmPassword}
              secureTextEntry
            />

            <Button
              title="Create Account"
              onPress={handleRegister}
              isLoading={isLoading}
              fullWidth
            />
          </View>

          <View className="mt-8 items-center">
            <Text className="text-gray-500">Already have an account?</Text>
            <Link href="/(auth)login" className="mt-2">
              <Text className="text-primary-600 font-medium">Sign In</Text>
            </Link>
          </View>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}

```

### app\(tabs)\_layout.tsx

```typescript
import { Tabs } from 'expo-router';
import { FontAwesome } from '@expo/vector-icons';

export default function TabLayout() {
  return (
    <Tabs
      screenOptions={{
        tabBarActiveTintColor: '#38B785',
        tabBarStyle: {
          borderTopWidth: 1,
          borderTopColor: '#E5E5E5',
          paddingBottom: 5,
          paddingTop: 5,
        },
      }}>
      <Tabs.Screen
        name="tasks"
        options={{
          title: 'Tasks',
          tabBarIcon: ({ color }) => (
            <FontAwesome name="tasks" size={24} color={color} />
          ),
        }}
      />
      <Tabs.Screen
        name="profile"
        options={{
          title: 'Profile',
          tabBarIcon: ({ color }) => (
            <FontAwesome name="user" size={24} color={color} />
          ),
        }}
      />
    </Tabs>
  );
}
```

### app\(tabs)\profile.tsx

```typescript
import React from 'react';
import { View, Text, Image, ScrollView, Pressable } from 'react-native';
import { FontAwesome } from '@expo/vector-icons';
import Animated, { 
  FadeInDown,
  useAnimatedStyle, 
  withSpring,
  useSharedValue
} from 'react-native-reanimated';
import { colors } from '@/constants/theme';

export default function ProfileScreen() {
  const scale = useSharedValue(1);
  
  const imageStyle = useAnimatedStyle(() => {
    return {
      transform: [{ scale: scale.value }]
    };
  });

  const handlePressIn = () => {
    scale.value = withSpring(1.1);
  };

  const handlePressOut = () => {
    scale.value = withSpring(1);
  };

  return (
    <ScrollView className="flex-1 bg-gray-50">
      {/* Profile Header */}
      <View className="pt-12 pb-6 px-6">
        <View className="items-center">
          <Pressable onPressIn={handlePressIn} onPressOut={handlePressOut}>
            <Animated.View 
              style={imageStyle} 
              className="w-24 h-24 rounded-full shadow-sm mb-4 overflow-hidden border-2 border-gray-100"
            >
              <Image
                source={{ uri: 'https://api.dicebear.com/7.x/avataaars/svg' }}
                className="w-full h-full"
              />
            </Animated.View>
          </Pressable>
          <Text className="text-xl font-semibold text-gray-800 mb-1">
            John Doe
          </Text>
          <Text className="text-gray-500 text-sm">
            john.doe@example.com
          </Text>
        </View>
      </View>

      {/* Statistics Cards */}
      <Animated.View 
        entering={FadeInDown.delay(200).springify()}
        className="flex-row justify-between mx-6 gap-3"
      >
        <View className="flex-1 bg-white p-4 rounded-xl border border-gray-100">
          <View className="items-center">
            <Text className="text-2xl font-semibold text-gray-800">12</Text>
            <Text className="text-gray-500 text-sm">Total</Text>
          </View>
        </View>
        
        <View className="flex-1 bg-white p-4 rounded-xl border border-gray-100">
          <View className="items-center">
            <Text className="text-2xl font-semibold text-gray-800">8</Text>
            <Text className="text-gray-500 text-sm">Done</Text>
          </View>
        </View>
        
        <View className="flex-1 bg-white p-4 rounded-xl border border-gray-100">
          <View className="items-center">
            <Text className="text-2xl font-semibold text-gray-800">4</Text>
            <Text className="text-gray-500 text-sm">Active</Text>
          </View>
        </View>
      </Animated.View>

      {/* Action Buttons */}
      <Animated.View 
        entering={FadeInDown.delay(400).springify()}
        className="mx-6 mt-6 space-y-3"
      >
        <Pressable className="flex-row items-center bg-white p-4 rounded-xl border border-gray-100">
          <View className="w-10 h-10 rounded-full items-center justify-center">
            <FontAwesome name="user" size={18} color={colors.gray[400]} />
          </View>
          <Text className="flex-1 ml-3 text-gray-700">Edit Profile</Text>
          <FontAwesome name="chevron-right" size={14} color={colors.gray[400]} />
        </Pressable>

        <Pressable className="flex-row items-center bg-white p-4 rounded-xl border border-gray-100">
          <View className="w-10 h-10 rounded-full items-center justify-center">
            <FontAwesome name="cog" size={18} color={colors.gray[400]} />
          </View>
          <Text className="flex-1 ml-3 text-gray-700">Settings</Text>
          <FontAwesome name="chevron-right" size={14} color={colors.gray[400]} />
        </Pressable>

        <Pressable className="flex-row items-center bg-white p-4 rounded-xl border border-gray-100 mb-8">
          <View className="w-10 h-10 rounded-full items-center justify-center">
            <FontAwesome name="sign-out" size={18} color={colors.gray[400]} />
          </View>
          <Text className="flex-1 ml-3 text-gray-700">Logout</Text>
          <FontAwesome name="chevron-right" size={14} color={colors.gray[400]} />
        </Pressable>
      </Animated.View>
    </ScrollView>
  );
}
```

### app\(tabs)\tasks.tsx

```typescript
import { View } from 'react-native';
import { TaskList } from '../../components/tasks/TaskList';
import { TaskForm } from '../../components/tasks/TaskForm';
import { useTasks } from '../../hooks/useTasks';
import { useAuth } from '../../context/AuthContext';
import type { CreateTaskPayload } from '../../types';

export default function TasksScreen() {
  const { tasks, isLoading, createTask, updateTask, deleteTask } = useTasks();
  const { user } = useAuth();

  const handleCreateTask = (taskData: CreateTaskPayload) => {
    if (user) {
      createTask({...taskData, userId: user.id});
    }
  };

  return (
    <View className="flex-1 bg-white shadow-sm">
      <View className="p-4">
        <TaskForm onSubmit={handleCreateTask} />
      </View>
      <TaskList
        tasks={tasks}
        isLoading={isLoading}
        onToggle={updateTask}
        onDelete={deleteTask}
      />
    </View>
  );
}
```

### app\+not-found.tsx

```typescript
import { Link, Stack } from 'expo-router';
import { StyleSheet } from 'react-native';

import { ThemedText } from '@/components/ThemedText';
import { ThemedView } from '@/components/ThemedView';

export default function NotFoundScreen() {
  return (
    <>
      <Stack.Screen options={{ title: 'Oops!' }} />
      <ThemedView style={styles.container}>
        <ThemedText type="title">This screen doesn't exist.</ThemedText>
        <Link href="/" style={styles.link}>
          <ThemedText type="link">Go to home screen!</ThemedText>
        </Link>
      </ThemedView>
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    padding: 20,
  },
  link: {
    marginTop: 15,
    paddingVertical: 15,
  },
});

```

### app\_layout.tsx

```typescript
import { Stack } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import '../styles/global.css';
import { AuthProvider } from '../context/AuthContext';

export default function RootLayout() {
  return (
    <AuthProvider>
      <GestureHandlerRootView style={{ flex: 1 }}>
        <StatusBar style="dark" />
        <Stack>
          <Stack.Screen name="(auth)" options={{ headerShown: false }} />
          <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
      </Stack>
    </GestureHandlerRootView>
    </AuthProvider>
  );
}

```

### app\index.tsx

```typescript
import { Redirect } from 'expo-router';
import { useAuth } from '../context/AuthContext';

export default function Index() {
  const { user } = useAuth();
  return <Redirect href={user ? "/(tabs)/tasks" : "/(auth)login"} />;
}
```

### components\Collapsible.tsx

```typescript
import { PropsWithChildren, useState } from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';

import { ThemedText } from '@/components/ThemedText';
import { ThemedView } from '@/components/ThemedView';
import { IconSymbol } from '@/components/ui/IconSymbol';
import { Colors } from '@/constants/Colors';
import { useColorScheme } from '@/hooks/useColorScheme';

export function Collapsible({ children, title }: PropsWithChildren & { title: string }) {
  const [isOpen, setIsOpen] = useState(false);
  const theme = useColorScheme() ?? 'light';

  return (
    <ThemedView>
      <TouchableOpacity
        style={styles.heading}
        onPress={() => setIsOpen((value) => !value)}
        activeOpacity={0.8}>
        <IconSymbol
          name="chevron.right"
          size={18}
          weight="medium"
          color={theme === 'light' ? Colors.light.icon : Colors.dark.icon}
          style={{ transform: [{ rotate: isOpen ? '90deg' : '0deg' }] }}
        />

        <ThemedText type="defaultSemiBold">{title}</ThemedText>
      </TouchableOpacity>
      {isOpen && <ThemedView style={styles.content}>{children}</ThemedView>}
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  heading: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 6,
  },
  content: {
    marginTop: 6,
    marginLeft: 24,
  },
});

```

### components\ExternalLink.tsx

```typescript
import { Link } from 'expo-router';
import { openBrowserAsync } from 'expo-web-browser';
import { type ComponentProps } from 'react';
import { Platform } from 'react-native';

type Props = Omit<ComponentProps<typeof Link>, 'href'> & { href: string };

export function ExternalLink({ href, ...rest }: Props) {
  return (
    <Link
      target="_blank"
      {...rest}
      href={href}
      onPress={async (event) => {
        if (Platform.OS !== 'web') {
          // Prevent the default behavior of linking to the default browser on native.
          event.preventDefault();
          // Open the link in an in-app browser.
          await openBrowserAsync(href);
        }
      }}
    />
  );
}

```

### components\HapticTab.tsx

```typescript
import { BottomTabBarButtonProps } from '@react-navigation/bottom-tabs';
import { PlatformPressable } from '@react-navigation/elements';
import * as Haptics from 'expo-haptics';

export function HapticTab(props: BottomTabBarButtonProps) {
  return (
    <PlatformPressable
      {...props}
      onPressIn={(ev) => {
        if (process.env.EXPO_OS === 'ios') {
          // Add a soft haptic feedback when pressing down on the tabs.
          Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light);
        }
        props.onPressIn?.(ev);
      }}
    />
  );
}

```

### components\HelloWave.tsx

```typescript
import { StyleSheet } from 'react-native';
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withTiming,
  withRepeat,
  withSequence,
} from 'react-native-reanimated';

import { ThemedText } from '@/components/ThemedText';

export function HelloWave() {
  const rotationAnimation = useSharedValue(0);

  rotationAnimation.value = withRepeat(
    withSequence(withTiming(25, { duration: 150 }), withTiming(0, { duration: 150 })),
    4 // Run the animation 4 times
  );

  const animatedStyle = useAnimatedStyle(() => ({
    transform: [{ rotate: `${rotationAnimation.value}deg` }],
  }));

  return (
    <Animated.View style={animatedStyle}>
      <ThemedText style={styles.text}>👋</ThemedText>
    </Animated.View>
  );
}

const styles = StyleSheet.create({
  text: {
    fontSize: 28,
    lineHeight: 32,
    marginTop: -6,
  },
});

```

### components\ParallaxScrollView.tsx

```typescript
import type { PropsWithChildren, ReactElement } from 'react';
import { StyleSheet } from 'react-native';
import Animated, {
  interpolate,
  useAnimatedRef,
  useAnimatedStyle,
  useScrollViewOffset,
} from 'react-native-reanimated';

import { ThemedView } from '@/components/ThemedView';
import { useBottomTabOverflow } from '@/components/ui/TabBarBackground';
import { useColorScheme } from '@/hooks/useColorScheme';

const HEADER_HEIGHT = 250;

type Props = PropsWithChildren<{
  headerImage: ReactElement;
  headerBackgroundColor: { dark: string; light: string };
}>;

export default function ParallaxScrollView({
  children,
  headerImage,
  headerBackgroundColor,
}: Props) {
  const colorScheme = useColorScheme() ?? 'light';
  const scrollRef = useAnimatedRef<Animated.ScrollView>();
  const scrollOffset = useScrollViewOffset(scrollRef);
  const bottom = useBottomTabOverflow();
  const headerAnimatedStyle = useAnimatedStyle(() => {
    return {
      transform: [
        {
          translateY: interpolate(
            scrollOffset.value,
            [-HEADER_HEIGHT, 0, HEADER_HEIGHT],
            [-HEADER_HEIGHT / 2, 0, HEADER_HEIGHT * 0.75]
          ),
        },
        {
          scale: interpolate(scrollOffset.value, [-HEADER_HEIGHT, 0, HEADER_HEIGHT], [2, 1, 1]),
        },
      ],
    };
  });

  return (
    <ThemedView style={styles.container}>
      <Animated.ScrollView
        ref={scrollRef}
        scrollEventThrottle={16}
        scrollIndicatorInsets={{ bottom }}
        contentContainerStyle={{ paddingBottom: bottom }}>
        <Animated.View
          style={[
            styles.header,
            { backgroundColor: headerBackgroundColor[colorScheme] },
            headerAnimatedStyle,
          ]}>
          {headerImage}
        </Animated.View>
        <ThemedView style={styles.content}>{children}</ThemedView>
      </Animated.ScrollView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  header: {
    height: HEADER_HEIGHT,
    overflow: 'hidden',
  },
  content: {
    flex: 1,
    padding: 32,
    gap: 16,
    overflow: 'hidden',
  },
});

```

### components\ThemedText.tsx

```typescript
import { Text, type TextProps, StyleSheet } from 'react-native';

import { useThemeColor } from '@/hooks/useThemeColor';

export type ThemedTextProps = TextProps & {
  lightColor?: string;
  darkColor?: string;
  type?: 'default' | 'title' | 'defaultSemiBold' | 'subtitle' | 'link';
};

export function ThemedText({
  style,
  lightColor,
  darkColor,
  type = 'default',
  ...rest
}: ThemedTextProps) {
  const color = useThemeColor({ light: lightColor, dark: darkColor }, 'text');

  return (
    <Text
      style={[
        { color },
        type === 'default' ? styles.default : undefined,
        type === 'title' ? styles.title : undefined,
        type === 'defaultSemiBold' ? styles.defaultSemiBold : undefined,
        type === 'subtitle' ? styles.subtitle : undefined,
        type === 'link' ? styles.link : undefined,
        style,
      ]}
      {...rest}
    />
  );
}

const styles = StyleSheet.create({
  default: {
    fontSize: 16,
    lineHeight: 24,
  },
  defaultSemiBold: {
    fontSize: 16,
    lineHeight: 24,
    fontWeight: '600',
  },
  title: {
    fontSize: 32,
    fontWeight: 'bold',
    lineHeight: 32,
  },
  subtitle: {
    fontSize: 20,
    fontWeight: 'bold',
  },
  link: {
    lineHeight: 30,
    fontSize: 16,
    color: '#0a7ea4',
  },
});

```

### components\ThemedView.tsx

```typescript
import { View, type ViewProps } from 'react-native';

import { useThemeColor } from '@/hooks/useThemeColor';

export type ThemedViewProps = ViewProps & {
  lightColor?: string;
  darkColor?: string;
};

export function ThemedView({ style, lightColor, darkColor, ...otherProps }: ThemedViewProps) {
  const backgroundColor = useThemeColor({ light: lightColor, dark: darkColor }, 'background');

  return <View style={[{ backgroundColor }, style]} {...otherProps} />;
}

```

### components\__tests__\ThemedText-test.tsx

```typescript
import * as React from 'react';
import renderer from 'react-test-renderer';

import { ThemedText } from '../ThemedText';

it(`renders correctly`, () => {
  const tree = renderer.create(<ThemedText>Snapshot test!</ThemedText>).toJSON();

  expect(tree).toMatchSnapshot();
});

```

### components\__tests__\__snapshots__\ThemedText-test.tsx.snap

```
// Jest Snapshot v1, https://goo.gl/fbAQLP

exports[`renders correctly 1`] = `
<Text
  style={
    [
      {
        "color": "#11181C",
      },
      {
        "fontSize": 16,
        "lineHeight": 24,
      },
      undefined,
      undefined,
      undefined,
      undefined,
      undefined,
    ]
  }
>
  Snapshot test!
</Text>
`;

```

### components\tasks\TaskForm.tsx

```typescript
// components/tasks/TaskForm.tsx
import React, { useState, useRef, useEffect } from 'react';
import { View, TextInput, Pressable, Text, Keyboard } from 'react-native';
import { FontAwesome } from '@expo/vector-icons';
import * as Haptics from 'expo-haptics';
import { BlurView } from 'expo-blur';
import Animated, {
  useAnimatedStyle,
  withSpring,
  withTiming,
  useSharedValue,
  interpolate,
  Extrapolation,
} from 'react-native-reanimated';
import { colors } from '@/constants/theme';
import { useAuth } from '@/context/AuthContext';
import type { CreateTaskPayload } from '@/types';

type TaskFormProps = {
  onSubmit: (task: CreateTaskPayload) => void;
  isLoading?: boolean;
  placeholder?: string;
  autoFocus?: boolean;
}

const AnimatedBlurView = Animated.createAnimatedComponent(BlurView);

export function TaskForm({ 
  onSubmit, 
  isLoading = false,
  placeholder = "Add a new task...",
  autoFocus = false,
}: TaskFormProps) {
  const [taskName, setTaskName] = useState('');
  const [isFocused, setIsFocused] = useState(false);
  const inputRef = useRef<TextInput>(null);
  const { user } = useAuth();
  
  const scale = useSharedValue(1);
  const blurIntensity = useSharedValue(0);
  const inputWidth = useSharedValue(1);

  useEffect(() => {
    if (autoFocus) {
      setTimeout(() => inputRef.current?.focus(), 100);
    }
  }, [autoFocus]);

  const handleSubmit = () => {
    if (taskName.trim() && !isLoading && user) {
      Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Medium);
      scale.value = withSpring(0.95, {}, () => {
        scale.value = withSpring(1);
      });
      onSubmit({
        name: taskName.trim(),
        done: false,
        userId: user.id // Include the user ID here
      });
      setTaskName('');
      Keyboard.dismiss();
    }
  };

  const handleFocus = () => {
    setIsFocused(true);
    blurIntensity.value = withTiming(100);
    inputWidth.value = withSpring(1.02);
  };

  const handleBlur = () => {
    setIsFocused(false);
    blurIntensity.value = withTiming(0);
    inputWidth.value = withSpring(1);
  };

  const containerStyle = useAnimatedStyle(() => ({
    transform: [{ scale: scale.value }],
  }));

  const inputContainerStyle = useAnimatedStyle(() => ({
    transform: [{ scale: inputWidth.value }],
  }));

  const blurStyle = useAnimatedStyle(() => ({
    opacity: interpolate(blurIntensity.value, [0, 100], [0, 1], Extrapolation.CLAMP),
  }));

  return (
    <Animated.View 
      style={{...containerStyle, borderRadius: 16}}
      className="px-4 pt-4 pb-2 bg-gray-50"
    >
      <View className="relative" style={{
        borderRadius: 16,
      }}>
        <AnimatedBlurView
          style={[blurStyle, { 
            position: 'absolute',
            top: -8,
            left: -8,
            right: -8,
            bottom: -8,
            borderRadius: 16,
          }]}
          intensity={20}
        />
        <Animated.View 
          style={{...inputContainerStyle, borderRadius: 16, overflow:"hidden"}}
          className="flex-row items-center bg-gray-50 rounded-2xl shadow-sm overflow-hidden border border-gray-100"
        >
          <View className="pl-4 pr-2"
          style={{
            paddingLeft: 10,
            paddingRight: 6,
          }}>
            <FontAwesome 
              name="pencil" 
              size={16} 
              color={isFocused ? colors.primary[500] : colors.gray[400]}
            />
          </View>

          <TextInput
            ref={inputRef}
            className="flex-1 h-12 pr-4 text-base text-gray-800 bg-gray-50"
            placeholder={placeholder}
            placeholderTextColor={colors.gray[400]}
            value={taskName}
            onChangeText={setTaskName}
            onSubmitEditing={handleSubmit}
            onFocus={handleFocus}
            onBlur={handleBlur}
            returnKeyType="done"
            maxLength={100}
            editable={!isLoading}
          />

          {taskName.length > 0 && (
            <Text className="text-xs text-gray-400 mr-2">
              {taskName.length}/100
            </Text>
          )}

          <Pressable
            onPress={handleSubmit}
            disabled={isLoading || !taskName.trim()}
            className="w-12 h-12 items-center justify-center"
            style={{ paddingRight: 10 }}
          >
            <View className="w-8 h-8 rounded-full items-center justify-center">
              {isLoading ? (
                <FontAwesome 
                  name="spinner" 
                  size={16} 
                  color={colors.gray[500]} 
                  className="animate-spin" 
                />
              ) : (
                <FontAwesome 
                  name="plus" 
                  size={16} 
                  color={taskName.trim() ? colors.gray[700] : colors.gray[500]} 
                />
              )}
            </View>
          </Pressable>
        </Animated.View>
      </View>

      {isFocused && (
        <Text className="text-xs text-gray-400 mt-2" style={{
          backgroundColor: "white"
        }}>
          Press enter or tap the plus button to add task
        </Text>
      )}
    </Animated.View>
  );
}
```

### components\tasks\TaskItem.tsx

```typescript
import React from 'react';
import { View, Text, Pressable } from 'react-native';
import { FontAwesome } from '@expo/vector-icons';
import * as Haptics from 'expo-haptics';
import type { Task } from '../../types';
import { colors } from '@/constants/theme';

interface TaskItemProps {
  task: Task;
  onToggle: (id: number, done: boolean) => void;
  onDelete: (id: number) => void;
}

export function TaskItem({ task, onToggle, onDelete }: TaskItemProps) {
  const handleToggle = () => {
    Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light);
    onToggle(task.id, !task.done);
  };

  const handleDelete = () => {
    Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Medium);
    onDelete(task.id);
  };

  return (
    <View  className="flex-row  items-center justify-between  border-b border-gray-200 " style={{
        gap: 12, paddingVertical: 5, paddingHorizontal: 10
    }}>
      <Pressable
        onPress={handleToggle}
        className="flex-row items-center flex-1 mr-4 gap-2 "
        style={{ flex: 1 , gap: 6 , backgroundColor: task.done? colors.primary[50] :colors.gray[50] , paddingVertical:10 , paddingHorizontal:10, borderRadius:12 }}
      >
        <View
          className={`w-6 h-6 rounded-full  mr-3 items-center justify-center flex  
            ${task.done ? 'bg-blue-500 border-blue-500' : 'border-gray-400'}`}
        >
          {task.done && (
            <FontAwesome name="check" size={14} color="green" />
          )}
        </View>
        <Text
          className={`text-base flex-1 ${
            task.done ? 'text-gray-400 line-through' : 'text-gray-800'
          }`}
        >
          {task.name}
        </Text>
      </Pressable>
      {!task.done && (
      <Pressable
        onPress={handleDelete}
        className="w-8 h-8 items-center justify-center"
      >
        <FontAwesome name="trash-o" size={20} color="#FF3B30" />
      </Pressable>
      )}
    </View>
  );
}
```

### components\tasks\TaskList.tsx

```typescript
import React from 'react';
import {
  View,
  Text,
  ScrollView,
  ActivityIndicator,
  RefreshControl,
  Dimensions,
  ViewStyle,
  StyleProp,
} from 'react-native';
import { TaskItem } from './TaskItem';
import type { Task } from '../../types';
import { FontAwesome } from '@expo/vector-icons';
import Animated, {
    AnimatedStyle,
  FadeInDown,
  FadeOut,
  LinearTransition,
  SlideInRight,
} from 'react-native-reanimated';
import { BlurView } from 'expo-blur';

const { width } = Dimensions.get('window');

interface TaskListProps {
  tasks: Task[];
  isLoading: boolean;
  onToggle: (id: number, done: boolean) => void;
  onDelete: (id: number) => void;
}

const LoadingState = () => (
  <View className="flex-1 items-center justify-center">
    <BlurView intensity={100} className="p-8 rounded-3xl">
      <ActivityIndicator size="large" color="#38B785" />
      <Text className="mt-4 text-gray-600 font-medium">Loading your tasks...</Text>
    </BlurView>
  </View>
);

const EmptyState = () => (
  <Animated.View 
    entering={FadeInDown.springify()}
    className="flex-1 items-center justify-center p-8"
  >
    <View className="w-32 h-32 bg-primary-100 rounded-full items-center justify-center mb-6">
      <FontAwesome name="tasks" size={48} color="#38B785" />
    </View>
    <Text className="text-2xl font-bold text-gray-800 text-center mb-2">
      No Tasks Yet
    </Text>
    <Text className="text-gray-500 text-base text-center leading-6 max-w-xs">
      Start adding tasks to organize your day. Your tasks will appear here.
    </Text>
    <View className="mt-8 items-center">
      <View className="flex-row items-center">
        <View className="w-10 h-10 bg-primary-100 rounded-full items-center justify-center mr-4">
          <FontAwesome name="plus" size={20} color="#38B785" />
        </View>
        <Text className="text-gray-600">
          Tap the button above to add a task
        </Text>
      </View>
    </View>
  </Animated.View>
);
interface TaskGroupProps {
    title: string;
    tasks: Task[];
    onToggle: (id: number, done: boolean) => void;
    onDelete: (id: number) => void;
    style?: StyleProp<AnimatedStyle<StyleProp<ViewStyle>>>
  }
const TaskGroup = ({ title, tasks, onToggle, onDelete, style }:TaskGroupProps) => (
  <Animated.View 
    entering={SlideInRight.springify()}
    layout={LinearTransition.springify()}
    className="mb-6 "
    style={style}
  >
    <Text className="text-gray-500 text-sm font-medium ml-4 mb-2">
      {title}
    </Text>
    <View className="bg-white rounded-2xl mx-4 my-4 shadow-sm overflow-hidden" style={{
      borderRadius: 16
    }}>
      {tasks.map((task, index) => (
        <Animated.View
          key={task.id}
          entering={FadeInDown.delay(index * 100).springify()}
          exiting={FadeOut}
          layout={LinearTransition.springify()}
        >
          <TaskItem
            task={task}
            onToggle={onToggle}
            onDelete={onDelete}
          />
          {index < tasks.length - 1 && (
            <View className="h-[1px] bg-gray-100 ml-16" />
          )}
        </Animated.View>
      ))}
    </View>
  </Animated.View>
);

export function TaskList({
  tasks,
  isLoading,
  onToggle,
  onDelete,
}: TaskListProps) {
  if (isLoading) {
    return <LoadingState />;
  }

  if (tasks.length === 0) {
    return <EmptyState />;
  }

  // Group tasks by completion status
  const pendingTasks = tasks.filter(task => !task.done);
  const completedTasks = tasks.filter(task => task.done);

  return (
    <ScrollView
      className="flex-1 bg-gray-50"
      refreshControl={
        <RefreshControl 
          refreshing={isLoading}
          tintColor="#38B785"
          colors={['#38B785']}
        />
      }
      showsVerticalScrollIndicator={false}
      contentContainerStyle={{ paddingVertical: 20, paddingHorizontal:10 }}
    >
      {pendingTasks.length > 0 && (
        <TaskGroup
          title={`PENDING (${pendingTasks.length})`}
          tasks={pendingTasks}
          onToggle={onToggle}
          onDelete={onDelete}
        />
      )}
      
      {completedTasks.length > 0 && (
        <TaskGroup
        style={{ marginTop: 20 }}
          title={`COMPLETED (${completedTasks.length})`}
          tasks={completedTasks}
          onToggle={onToggle}
          onDelete={onDelete}
        />
      )}

      {/* Bottom padding */}
      <View className="h-20" />
    </ScrollView>
  );
}
```

### components\ui\Button.tsx

```typescript
import React from 'react';
import { Pressable, Text, ActivityIndicator } from 'react-native';
import * as Haptics from 'expo-haptics';

interface ButtonProps {
  onPress: () => void;
  title: string;
  variant?: 'primary' | 'secondary' | 'danger';
  isLoading?: boolean;
  disabled?: boolean;
  fullWidth?: boolean;
}

export function Button({
  onPress,
  title,
  variant = 'primary',
  isLoading = false,
  disabled = false,
  fullWidth = false,
}: ButtonProps) {
  const handlePress = () => {
    Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light);
    onPress();
  };

  const getVariantStyles = () => {
    switch (variant) {
      case 'primary':
        return 'h-14 bg-blue-500 active:bg-blue-600 disabled:bg-blue-300';
      case 'secondary':
        return 'h-14 bg-gray-200 active:bg-gray-300 disabled:bg-gray-100';
      case 'danger':
        return 'h-14 bg-red-500 active:bg-red-600 disabled:bg-red-300';
      default:
        return 'h-14 bg-blue-500 active:bg-blue-600 disabled:bg-blue-300';
    }
  };

  const getTextColor = () => {
    switch (variant) {
      case 'secondary':
        return 'text-gray-800';
      default:
        return 'text-white';
    }
  };

  return (
    <Pressable
      onPress={handlePress}
      disabled={disabled || isLoading}
      className={`
        h-14 px-6 rounded-lg items-center justify-center
        ${getVariantStyles()}
        ${fullWidth ? 'w-full' : 'w-auto'}
      `}
    >
      {isLoading ? (
        <ActivityIndicator color={variant === 'secondary' ? '#1F2937' : 'white'} />
      ) : (
        <Text
        style={{ height: 32 }}
          className={`
            font-medium text-base h-14 text-center flex justify-center items-center
            ${getTextColor()}
          `}
        >
          {title}
        </Text>
      )}
    </Pressable>
  );
}
```

### components\ui\IconSymbol.ios.tsx

```typescript
import { SymbolView, SymbolViewProps, SymbolWeight } from 'expo-symbols';
import { StyleProp, ViewStyle } from 'react-native';

export function IconSymbol({
  name,
  size = 24,
  color,
  style,
  weight = 'regular',
}: {
  name: SymbolViewProps['name'];
  size?: number;
  color: string;
  style?: StyleProp<ViewStyle>;
  weight?: SymbolWeight;
}) {
  return (
    <SymbolView
      weight={weight}
      tintColor={color}
      resizeMode="scaleAspectFit"
      name={name}
      style={[
        {
          width: size,
          height: size,
        },
        style,
      ]}
    />
  );
}

```

### components\ui\IconSymbol.tsx

```typescript
// This file is a fallback for using MaterialIcons on Android and web.

import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import { SymbolWeight } from 'expo-symbols';
import React from 'react';
import { OpaqueColorValue, StyleProp, ViewStyle } from 'react-native';

// Add your SFSymbol to MaterialIcons mappings here.
const MAPPING = {
  // See MaterialIcons here: https://icons.expo.fyi
  // See SF Symbols in the SF Symbols app on Mac.
  'house.fill': 'home',
  'paperplane.fill': 'send',
  'chevron.left.forwardslash.chevron.right': 'code',
  'chevron.right': 'chevron-right',
} as Partial<
  Record<
    import('expo-symbols').SymbolViewProps['name'],
    React.ComponentProps<typeof MaterialIcons>['name']
  >
>;

export type IconSymbolName = keyof typeof MAPPING;

/**
 * An icon component that uses native SFSymbols on iOS, and MaterialIcons on Android and web. This ensures a consistent look across platforms, and optimal resource usage.
 *
 * Icon `name`s are based on SFSymbols and require manual mapping to MaterialIcons.
 */
export function IconSymbol({
  name,
  size = 24,
  color,
  style,
}: {
  name: IconSymbolName;
  size?: number;
  color: string | OpaqueColorValue;
  style?: StyleProp<ViewStyle>;
  weight?: SymbolWeight;
}) {
  return <MaterialIcons color={color} size={size} name={MAPPING[name]} style={style} />;
}

```

### components\ui\Input.tsx

```typescript
import React from 'react';
import { View, Text, TextInput, TextInputProps } from 'react-native';

interface InputProps extends TextInputProps {
  label?: string;
  error?: string;
}

export function Input({ label, error, ...props }: InputProps) {
  return (
    <View className="mb-4">
      {label && (
        <Text className="mb-1 text-sm font-medium text-gray-700">
          {label}
        </Text>
      )}
      <TextInput
        className={`
          h-12 px-4 bg-white border rounded-lg text-base
          ${error ? 'border-red-500' : 'border-gray-300'}
          ${props.editable === false ? 'bg-gray-100' : ''}
        `}
        placeholderTextColor="#9CA3AF"
        {...props}
      />
      {error && (
        <Text className="mt-1 text-sm text-red-500">
          {error}
        </Text>
      )}
    </View>
  );
}
```

### components\ui\TabBarBackground.ios.tsx

```typescript
import { useBottomTabBarHeight } from '@react-navigation/bottom-tabs';
import { BlurView } from 'expo-blur';
import { StyleSheet } from 'react-native';
import { useSafeAreaInsets } from 'react-native-safe-area-context';

export default function BlurTabBarBackground() {
  return (
    <BlurView
      // System chrome material automatically adapts to the system's theme
      // and matches the native tab bar appearance on iOS.
      tint="systemChromeMaterial"
      intensity={100}
      style={StyleSheet.absoluteFill}
    />
  );
}

export function useBottomTabOverflow() {
  const tabHeight = useBottomTabBarHeight();
  const { bottom } = useSafeAreaInsets();
  return tabHeight - bottom;
}

```

### components\ui\TabBarBackground.tsx

```typescript
// This is a shim for web and Android where the tab bar is generally opaque.
export default undefined;

export function useBottomTabOverflow() {
  return 0;
}

```

### constants\Colors.ts

```typescript
/**
 * Below are the colors that are used in the app. The colors are defined in the light and dark mode.
 * There are many other ways to style your app. For example, [Nativewind](https://www.nativewind.dev/), [Tamagui](https://tamagui.dev/), [unistyles](https://reactnativeunistyles.vercel.app), etc.
 */

const tintColorLight = '#0a7ea4';
const tintColorDark = '#fff';

export const Colors = {
  light: {
    text: '#11181C',
    background: '#fff',
    tint: tintColorLight,
    icon: '#687076',
    tabIconDefault: '#687076',
    tabIconSelected: tintColorLight,
  },
  dark: {
    text: '#ECEDEE',
    background: '#151718',
    tint: tintColorDark,
    icon: '#9BA1A6',
    tabIconDefault: '#9BA1A6',
    tabIconSelected: tintColorDark,
  },
};

```

### constants\theme.ts

```typescript
export const colors = {
    primary: {
        50: '#F0FFF4',
        100: '#E6FFEE',
        200: '#C3F9DB',
        300: '#9EEDC8',
        400: '#65D6A3',
        500: '#38B785',
        600: '#2D9D71',
        700: '#25805C',
        800: '#1E654A',
        900: '#164B38',
    },
    gray: {
      50: '#F9FAFB',
      100: '#F3F4F6',
      200: '#E5E7EB',
      300: '#D1D5DB',
      400: '#9CA3AF',
      500: '#6B7280',
      600: '#4B5563',
      700: '#374151',
      800: '#1F2937',
      900: '#111827',
    },
    success: {
      50: '#F0FDF4',
      500: '#22C55E',
      600: '#16A34A',
    },
    error: {
      50: '#FEF2F2',
      500: '#EF4444',
      600: '#DC2626',
    },
  } as const
  
  export const spacing = {
    0: 0,
    1: 4,
    2: 8,
    3: 12,
    4: 16,
    5: 20,
    6: 24,
    8: 32,
    10: 40,
    12: 48,
    16: 64,
    20: 80,
    24: 96,
  };
  
  export const fontSize = {
    xs: 12,
    sm: 14,
    base: 16,
    lg: 18,
    xl: 20,
    '2xl': 24,
    '3xl': 30,
    '4xl': 36,
  };
  
  export const borderRadius = {
    none: 0,
    sm: 2,
    base: 4,
    md: 6,
    lg: 8,
    xl: 12,
    '2xl': 16,
    full: 9999,
  };
```

### context\AuthContext.tsx

```typescript
import React, { createContext, useState, useContext, useEffect } from 'react';
import { router, useSegments } from 'expo-router';
import AuthService from '../services/tasks';
import * as SecureStore from 'expo-secure-store';

interface User {
  id: number;
  email: string;
}

interface AuthContextType {
  user: User | null;
  isLoading: boolean;
  signIn: (email: string, password: string) => Promise<void>;
  signUp: (email: string, password: string) => Promise<void>;
  signOut: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | null>(null);

// This hook can be used to access the user info.
export function useAuth() {
  return useContext(AuthContext) as AuthContextType;
}

// This hook will protect the route access based on user authentication.
function useProtectedRoute(user: User | null) {
  const segments = useSegments();

  useEffect(() => {
    const inAuthGroup = segments[0] === '(auth)';

    if (!user && !inAuthGroup) {
      // If the user is not signed in and the initial segment is not in the auth group.
      router.replace('/(auth)login');
    } else if (user && inAuthGroup) {
      // Redirect away from auth group pages if the user is signed in.
      router.replace('/(tabs)/tasks');
    }
  }, [user, segments]);
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useProtectedRoute(user);

  useEffect(() => {
    // Check for stored authentication token on app load
    loadStoredAuth();
  }, []);

  async function loadStoredAuth() {
    try {
      const storedUser = await SecureStore.getItemAsync('user');
      if (storedUser) {
        setUser(JSON.parse(storedUser));
      }
    } catch (error) {
      console.error('Error loading stored auth:', error);
    } finally {
      setIsLoading(false);
    }
  }

  async function signIn(email: string, password: string) {
    try {
      setIsLoading(true);
      const response = await AuthService.login({ email, password });
      await SecureStore.setItemAsync('user', JSON.stringify(response.user));
      setUser(response.user);
    } finally {
      setIsLoading(false);
    }
  }

  async function signUp(email: string, password: string) {
    try {
      setIsLoading(true);
      const user = await AuthService.register({ email, password });
      // After registration, automatically sign in
      await signIn(email, password);
    } finally {
      setIsLoading(false);
    }
  }

  async function signOut() {
    try {
      setIsLoading(true);
      await AuthService.logout();
      await SecureStore.deleteItemAsync('user');
      setUser(null);
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <AuthContext.Provider value={{ user, isLoading, signIn, signUp, signOut }}>
      {children}
    </AuthContext.Provider>
  );
}

```

### hooks\useColorScheme.ts

```typescript
export { useColorScheme } from 'react-native';

```

### hooks\useColorScheme.web.ts

```typescript
import { useEffect, useState } from 'react';
import { useColorScheme as useRNColorScheme } from 'react-native';

/**
 * To support static rendering, this value needs to be re-calculated on the client side for web
 */
export function useColorScheme() {
  const [hasHydrated, setHasHydrated] = useState(false);

  useEffect(() => {
    setHasHydrated(true);
  }, []);

  const colorScheme = useRNColorScheme();

  if (hasHydrated) {
    return colorScheme;
  }

  return 'light';
}

```

### hooks\useTasks.ts

```typescript
import { useState, useEffect, useCallback } from 'react';
import { APIError ,TasksService } from '../services/tasks';
import type { CreateTaskPayload, Task } from '../types';

export function useTasks() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchTasks = useCallback(async () => {
    try {
      setIsLoading(true);
      const fetchedTasks = await TasksService.getTasks();
      setTasks(fetchedTasks);
      setError(null);
    } catch (err) {
      if (err instanceof APIError) {
        setError(err.message);
      } else {
        setError('Failed to fetch tasks');
      }
    } finally {
      setIsLoading(false);
    }
  }, []);

  const createTask = useCallback(async (task: CreateTaskPayload) => {
    try {
      const newTask = await TasksService.createTask({...task , done: false});
      setTasks((prev) => [...prev, newTask]);
      setError(null);
    } catch (err) {
      if (err instanceof APIError) {
        setError(err.message);
      } else {
        setError('Failed to create task');
      }
    }
  }, []);

  const updateTask = useCallback(async (id: number, done: boolean) => {
    try {
      const updatedTask = await TasksService.updateTask(id, { done });
      setTasks((prev) =>
        prev.map((task) =>
          task.id === id ? { ...task, ...updatedTask } : task
        )
      );
      setError(null);
    } catch (err) {
      if (err instanceof APIError) {
        setError(err.message);
      } else {
        setError('Failed to update task');
      }
    }
  }, []);

  const deleteTask = useCallback(async (id: number) => {
    try {
      await TasksService.deleteTask(id);
      setTasks((prev) => prev.filter((task) => task.id !== id));
      setError(null);
    } catch (err) {
      if (err instanceof APIError) {
        setError(err.message);
      } else {
        setError('Failed to delete task');
      }
    }
  }, []);

  useEffect(() => {
    fetchTasks();
  }, [fetchTasks]);

  return {
    tasks,
    isLoading,
    error,
    createTask,
    updateTask,
    deleteTask,
    refetch: fetchTasks,
  };
}
```

### hooks\useThemeColor.ts

```typescript
/**
 * Learn more about light and dark modes:
 * https://docs.expo.dev/guides/color-schemes/
 */

import { Colors } from '@/constants/Colors';
import { useColorScheme } from '@/hooks/useColorScheme';

export function useThemeColor(
  props: { light?: string; dark?: string },
  colorName: keyof typeof Colors.light & keyof typeof Colors.dark
) {
  const theme = useColorScheme() ?? 'light';
  const colorFromProps = props[theme];

  if (colorFromProps) {
    return colorFromProps;
  } else {
    return Colors[theme][colorName];
  }
}

```

### scripts\reset-project.js

```javascript
#!/usr/bin/env node

/**
 * This script is used to reset the project to a blank state.
 * It moves the /app, /components, /hooks, /scripts, and /constants directories to /app-example and creates a new /app directory with an index.tsx and _layout.tsx file.
 * You can remove the `reset-project` script from package.json and safely delete this file after running it.
 */

const fs = require("fs");
const path = require("path");

const root = process.cwd();
const oldDirs = ["app", "components", "hooks", "constants", "scripts"];
const newDir = "app-example";
const newAppDir = "app";
const newDirPath = path.join(root, newDir);

const indexContent = `import { Text, View } from "react-native";

export default function Index() {
  return (
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text>Edit app/index.tsx to edit this screen.</Text>
    </View>
  );
}
`;

const layoutContent = `import { Stack } from "expo-router";

export default function RootLayout() {
  return <Stack />;
}
`;

const moveDirectories = async () => {
  try {
    // Create the app-example directory
    await fs.promises.mkdir(newDirPath, { recursive: true });
    console.log(`📁 /${newDir} directory created.`);

    // Move old directories to new app-example directory
    for (const dir of oldDirs) {
      const oldDirPath = path.join(root, dir);
      const newDirPath = path.join(root, newDir, dir);
      if (fs.existsSync(oldDirPath)) {
        await fs.promises.rename(oldDirPath, newDirPath);
        console.log(`➡️ /${dir} moved to /${newDir}/${dir}.`);
      } else {
        console.log(`➡️ /${dir} does not exist, skipping.`);
      }
    }

    // Create new /app directory
    const newAppDirPath = path.join(root, newAppDir);
    await fs.promises.mkdir(newAppDirPath, { recursive: true });
    console.log("\n📁 New /app directory created.");

    // Create index.tsx
    const indexPath = path.join(newAppDirPath, "index.tsx");
    await fs.promises.writeFile(indexPath, indexContent);
    console.log("📄 app/index.tsx created.");

    // Create _layout.tsx
    const layoutPath = path.join(newAppDirPath, "_layout.tsx");
    await fs.promises.writeFile(layoutPath, layoutContent);
    console.log("📄 app/_layout.tsx created.");

    console.log("\n✅ Project reset complete. Next steps:");
    console.log(
      "1. Run `npx expo start` to start a development server.\n2. Edit app/index.tsx to edit the main screen.\n3. Delete the /app-example directory when you're done referencing it."
    );
  } catch (error) {
    console.error(`Error during script execution: ${error}`);
  }
};

moveDirectories();

```

### services\tasks.ts

```typescript
import ky, { HTTPError } from 'ky';

const API_URL = 'https://hono-starter-cloudflare-2.aissaouimoncefdev.workers.dev';

// Type definitions
type User = {
    id: number;
    email: string;
    createdAt: string | null;
    updatedAt: string | null;
}

type Task = {
    id: number;
    userId: number;
    name: string;
    done: boolean;
    createdAt: string | null;
    updatedAt: string | null;
}

type LoginResponse = {
    token: string;
    user: User;
}

type RegisterPayload = {
    email: string;
    password: string;
}

type LoginPayload = {
    email: string;
    password: string;
}

type CreateTaskPayload = {
    userId: number;
    name: string;
    done: boolean;
}

type UpdateTaskPayload = {
    userId?: number;
    name?: string;
    done?: boolean;
}

type ValidationError = {
    success: false;
    error: {
        issues: Array<{
            code: string;
            path: Array<string | number>;
            message?: string;
        }>;
        name: string;
    };
}

interface APIErrorResponse {
    message: string;
}

// API Client configuration
let authToken: string | null = null;

const api = ky.create({
    prefixUrl: API_URL,
    timeout: 30000,
    hooks: {
        beforeRequest: [
            request => {
                if (authToken) {
                    request.headers.set('Authorization', `Bearer ${authToken}`);
                }
            }
        ],
        beforeError: [
            async (error: HTTPError & { response?: Response }) => {
                const { response } = error;
                if (response?.body) {
                    try {
                        const errorData = await response.json() as APIErrorResponse;
                        error.name = 'APIError';
                        error.message = errorData.message || 'Unknown error occurred';
                    } catch (e) {
                        error.message = 'Failed to parse error response';
                    }
                }
                return error;
            },
        ],
    },
});

// Auth Service
export class AuthService {
    // Register a new user
    static async register(payload: RegisterPayload): Promise<User> {
        return api.post('auth/register', {
            json: payload,
        }).json<User>();
    }

    // Login user
    static async login(payload: LoginPayload): Promise<LoginResponse> {
        const response = await api.post('auth/login', {
            json: payload,
        }).json<LoginResponse>();
        
        // Store the token for subsequent requests
        authToken = response.token;
        
        return response;
    }

    // Logout user
    static async logout(): Promise<void> {
        await api.post('auth/logout');
        authToken = null;
    }
}

// Tasks API Service
export class TasksService {
    // Get all tasks
    static async getTasks(): Promise<Task[]> {
        return api.get('tasks').json<Task[]>();
    }

    // Get a single task by ID
    static async getTask(id: number): Promise<Task> {
        return api.get(`tasks/${id}`).json<Task>();
    }

    // Create a new task
    static async createTask(payload: CreateTaskPayload): Promise<Task> {
        return api.post('tasks', {
            json: payload,
        }).json<Task>();
    }

    // Update an existing task
    static async updateTask(id: number, payload: UpdateTaskPayload): Promise<Task> {
        return api.patch(`tasks/${id}`, {
            json: payload,
        }).json<Task>();
    }

    // Delete a task
    static async deleteTask(id: number): Promise<void> {
        await api.delete(`tasks/${id}`);
    }
}

// Error handling utility
export class APIError extends Error {
    constructor(message: string) {
        super(message);
        this.name = 'APIError';
    }

    static isValidationError(error: unknown): error is ValidationError {
        return (
            typeof error === 'object' &&
            error !== null &&
            'success' in error &&
            error.success === false &&
            'error' in error &&
            typeof error.error === 'object' &&
            error.error !== null &&
            'issues' in error.error &&
            Array.isArray((error as ValidationError).error.issues)
        );
    }
}

// Usage example:
/*
try {
    // Register a new user
    const user = await AuthService.register({
        email: "user@example.com",
        password: "password123"
    });

    // Login
    const { token, user: loggedInUser } = await AuthService.login({
        email: "user@example.com",
        password: "password123"
    });

    // Get all tasks
    const tasks = await TasksService.getTasks();

    // Create a new task
    const newTask = await TasksService.createTask({
        userId: loggedInUser.id,
        name: "Complete project",
        done: false
    });

    // Update a task
    const updatedTask = await TasksService.updateTask(newTask.id, {
        done: true
    });

    // Delete a task
    await TasksService.deleteTask(newTask.id);

    // Logout
    await AuthService.logout();

} catch (error) {
    if (APIError.isValidationError(error)) {
        // Handle validation errors
        console.error('Validation error:', error.error.issues);
    } else if (error instanceof APIError) {
        // Handle API errors
        console.error('API error:', error.message);
    } else {
        // Handle other errors
        console.error('Unknown error:', error);
    }
}
*/

export { AuthService as default };
```

### styles\global.css

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

### types\index.ts

```typescript
export interface Task {
  id: number;
  userId: number;  // Added userId
  name: string;
  done: boolean;
  createdAt: string | null;
  updatedAt: string | null;
}

export interface CreateTaskPayload {
  userId: number;  // Added userId
  name: string;
  done: boolean;
}

export interface UpdateTaskPayload {
  name?: string;
  done?: boolean;
}
  export interface ValidationError {
    success: false;
    error: {
      issues: Array<{
        code: string;
        path: Array<string | number>;
        message?: string;
      }>;
      name: string;
    };
  }
  
  export interface APIErrorResponse {
    message: string;
  }
```

