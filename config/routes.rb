Rails.application.routes.draw do

  devise_for :users, controllers: { registrations: 'registrations' } 

  devise_for :admins

  devise_scope :user do
    get 'sign_up', to: 'devise/registrations#new'
    get 'sign_in', to: 'devise/sessions#create'
    get 'sign_out', to: 'devise/sessions#destroy'

    unauthenticated :user do
      root to: 'devise/registrations#new', as: :unauthenticated_root
    end
  end

  authenticated :user do
    root to: 'homepages#index', as: :authenticated_root
  end

  root 'homepages#homepage'

  resources :after_signups

  resources :chat
  resources :users
  resources :games


  resource :homepage, only: [:index] do
    get :homepage
  end
  
  namespace :user do
    resources :account_settings, only: [] do
      collection do
        get :edit_personal_info
        get :edit_profile_info

        put :update_personal_info
        put :update_profile_info
      end
    end
  end

  # You can have the root of your site routed with "root"
  # root 'welcome#index'

  # Example of regular route:
  #   get 'products/:id' => 'catalog#view'

  # Example of named route that can be invoked with purchase_url(id: product.id)
  #   get 'products/:id/purchase' => 'catalog#purchase', as: :purchase

  # Example resource route (maps HTTP verbs to controller actions automatically):
  #   resources :products

  # Example resource route with options:
  #   resources :products do
  #     member do
  #       get 'short'
  #       post 'toggle'
  #     end
  #
  #     collection do
  #       get 'sold'
  #     end
  #   end

  # Example resource route with sub-resources:
  #   resources :products do
  #     resources :comments, :sales
  #     resource :seller
  #   end

  # Example resource route with more complex sub-resources:
  #   resources :products do
  #     resources :comments
  #     resources :sales do
  #       get 'recent', on: :collection
  #     end
  #   end

  # Example resource route with concerns:
  #   concern :toggleable do
  #     post 'toggle'
  #   end
  #   resources :posts, concerns: :toggleable
  #   resources :photos, concerns: :toggleable

  # Example resource route within a namespace:
  #   namespace :admin do
  #     # Directs /admin/products/* to Admin::ProductsController
  #     # (app/controllers/admin/products_controller.rb)
  #     resources :products
  #   end
end